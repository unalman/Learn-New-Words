package main

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

var w = &Word{
	Id:               2,
	Main_language:    "main",
	Foreign_language: "ana",
	Created_at:       time.Now().UTC(),
}
var ws = []Word{
	{
		Id:               1,
		Main_language:    "main",
		Foreign_language: "ana",
		Created_at:       time.Now().UTC(),
	},
	{
		Id:               2,
		Main_language:    "slow",
		Foreign_language: "yavaş",
		Created_at:       time.Now().UTC(),
	},
}

func MockCreatePostgreSqlConnection() (server *PostgreSqlInfo, mock sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return &PostgreSqlInfo{
		db: db,
	}, mock
}
func TestCreateWord(t *testing.T) {
	s, mock := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(w.Id)

	query := "INSERT INTO word (main_language, foreign_language, created_at) VALUES ($1, $2, $3) RETURNING id"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(w.Main_language, w.Foreign_language, AnyTime{}).WillReturnRows(rows)

	word, err := s.CreateWord(w.Main_language, w.Foreign_language)
	assert.NotNil(t, word)
	assert.Equal(t, w.Id, word.Id)
	assert.NoError(t, err)
}
func TestCreateWordFailed(t *testing.T) {
	s, mock := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()

	query := "INSERT INTO word (main_language, foreign_language, created_at) VALUES ($1, $2, $3) RETURNING id"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(w.Main_language, w.Foreign_language, AnyTime{}).WillReturnError(sql.ErrConnDone)

	word, err := s.CreateWord(w.Main_language, w.Foreign_language)
	assert.Nil(t, word)
	assert.Error(t, err)
	if !errors.Is(err, sql.ErrConnDone) {
		t.Errorf("[CreateWord] word error assertion failed, got err %v", err)
	}
}
func TestGetWordById(t *testing.T) {
	s, mock := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()
	rows := sqlmock.NewRows([]string{"id", "created_at", "main_language", "foreign_language"}).AddRow(w.Id, w.Created_at, w.Main_language, w.Foreign_language)
	query := "SELECT * FROM word WHERE id = $1"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(w.Id).WillReturnRows(rows)

	word, err := s.GetWordById(w.Id)
	assert.NotNil(t, word)
	assert.NoError(t, err)
}
func TestGetWordByIdFailed(t *testing.T) {
	testCases := []struct {
		name             string
		expectedResponse *Word
		expectedErr      error
	}{
		{
			name:             "TestGetWordByIdFailedSQL",
			expectedResponse: nil,
			expectedErr:      sql.ErrConnDone,
		},
		{
			name:             "TestGetWordByIdFailedById",
			expectedResponse: nil,
			expectedErr:      fmt.Errorf("id %s does not exist", strconv.Itoa(w.Id)),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s, mock := MockCreatePostgreSqlConnection()
			defer func() {
				s.db.Close()
			}()

			query := "SELECT * FROM word WHERE id = $1"
			mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(w.Id).WillReturnError(tc.expectedErr)

			word, err := s.GetWordById(w.Id)
			assert.Nil(t, word)
			assert.Error(t, err)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("[GetWordById] word error assertion failed, got err %v", err)
			}
		})
	}
}
func TestGetListWord(t *testing.T) {
	s, mock := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()
	rows := sqlmock.NewRows([]string{"id", "created_at", "main_language", "foreign_language"})
	var expectedLength = len(ws)
	for _, wm := range ws {
		rows.AddRow(wm.Id, wm.Created_at, wm.Main_language, wm.Foreign_language)
	}
	query := "SELECT * FROM word"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	words, err := s.GetListWord()
	assert.NotNil(t, words)
	assert.NoError(t, err)
	assert.Len(t, words, expectedLength)
}
func TestGetListWordFailed(t *testing.T) {
	s, mock := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()

	query := "SELECT * FROM word"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(sql.ErrConnDone)
	words, err := s.GetListWord()
	assert.Nil(t, words)
	assert.Error(t, err)
	if !errors.Is(err, sql.ErrConnDone) {
		t.Errorf("[TestGetListWord] word error assertion failed, got err %v", err)
	}
}
func TestUpdateWord(t *testing.T) {
	s, mock := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()
	rows := sqlmock.NewRows([]string{"id", "created_at", "main_language", "foreign_language"})
	query := "UPDATE word SET main_language=$1, foreign_language=$2 WHERE id=$3"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(w.Main_language, w.Foreign_language, w.Id).WillReturnRows(rows)
	err := s.UpdateWord(w)
	assert.NoError(t, err)
}
func TestUpdateWordFailed(t *testing.T) {
	s, mock := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()
	query := "UPDATE word SET main_language=$1, foreign_language=$2 WHERE id=$3"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(w.Main_language, w.Foreign_language, w.Id).WillReturnError(sql.ErrConnDone)
	err := s.UpdateWord(w)
	assert.Error(t, err)
}
func TestDeleteWord(t *testing.T) {
	s, mock := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()
	query := "DELETE FROM word WHERE id=$1"
	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(w.Id).WillReturnResult(sqlmock.NewResult(0, 1))
	err := s.DeleteWord(w.Id)
	assert.NoError(t, err)
}
func TestDeleteWordFailed(t *testing.T) {
	s, mock := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()
	query := "DELETE FROM word WHERE id=$1"
	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(w.Id).WillReturnResult(sqlmock.NewResult(0, 0))
	err := s.DeleteWord(w.Id)
	assert.Error(t, err)
}
func TestExistsWord(t *testing.T) {
	s, mock := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()
	row := sqlmock.NewRows([]string{"exists"}).AddRow(true)
	query := "SELECT exists (SELECT 1 FROM word WHERE id=$1 LIMIT 1)"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(w.Id).WillReturnRows(row)

	isExists, err := s.ExistsWord(w.Id)
	assert.NotNil(t, isExists)
	assert.True(t, isExists)
	assert.NoError(t, err)
}
func TestExistsWordFailed(t *testing.T) {
	s, mock := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()
	query := "SELECT exists (SELECT 1 FROM word WHERE id=$1 LIMIT 1)"
	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(w.Id).WillReturnResult(sqlmock.NewResult(0, 0))

	isExists, err := s.ExistsWord(w.Id)
	assert.Error(t, err)
	assert.False(t, isExists)
}
