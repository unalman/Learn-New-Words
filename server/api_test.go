package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var wordModel = &Word{
	Id:               2,
	Main_language:    "main",
	Foreign_language: "ana",
	Created_at:       time.Now().UTC(),
}
var newWord = Word{
	Main_language:    wordModel.Main_language,
	Foreign_language: wordModel.Foreign_language,
}
var wordJson, _ = json.Marshal(newWord)

func TestApiGetWords(t *testing.T) {
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

	server := NewApiServer(os.Getenv("API_HOST")+":"+os.Getenv("API_PORT"), s)
	router := server.SetupRouter()
	router = GetWords(server, router)
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	req, _ := http.NewRequest("GET", "/words", nil)
	req.Header.Add("Accept", "application/json")
	router.ServeHTTP(w, req)
	var resM ResponseModel[[]*Word]
	decoder := json.NewDecoder(w.Body)

	err := decoder.Decode(&resM)
	if err != nil {
		t.Errorf("Json decoder doesn't work: %s", err)
	}
	if resM.Data == nil {
		t.Errorf("Response is nil")
	}

	assert.Equal(t, 200, w.Code)
	assert.Len(t, resM.Data, expectedLength)
}
func TestApiGetWordsFailed(t *testing.T) {
	s, mock := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()

	query := "SELECT * FROM word"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(sql.ErrConnDone)

	server := NewApiServer(os.Getenv("API_HOST")+":"+os.Getenv("API_PORT"), s)
	router := server.SetupRouter()
	router = GetWords(server, router)
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	req, _ := http.NewRequest("GET", "/words", nil)
	req.Header.Add("Accept", "application/json")
	router.ServeHTTP(w, req)
	var resM ResponseModel[[]*Word]
	decoder := json.NewDecoder(w.Body)

	err := decoder.Decode(&resM)
	if err != nil {
		t.Errorf("Json decoder doesn't work: %s", err)
	}

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, sql.ErrConnDone.Error(), resM.Error)
}
func TestApiGetRandomWords(t *testing.T) {
	s, mock := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()
	rows := sqlmock.NewRows([]string{"id", "created_at", "main_language", "foreign_language"})
	var expectedLength = len(ws)
	for _, wm := range ws {
		rows.AddRow(wm.Id, wm.Created_at, wm.Main_language, wm.Foreign_language)
	}
	query := "SELECT * FROM word ORDER BY RANDOM() LIMIT 3"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	server := NewApiServer(os.Getenv("API_HOST")+":"+os.Getenv("API_PORT"), s)
	router := server.SetupRouter()
	router = GetRandomWords(server, router)
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	req, _ := http.NewRequest("GET", "/random-words", nil)
	req.Header.Add("Accept", "application/json")
	router.ServeHTTP(w, req)
	var resM ResponseModel[[]*Word]
	decoder := json.NewDecoder(w.Body)

	err := decoder.Decode(&resM)
	if err != nil {
		t.Errorf("Json decoder doesn't work: %s", err)
	}
	if resM.Data == nil {
		t.Errorf("Response is nil")
	}

	assert.Equal(t, 200, w.Code)
	assert.Len(t, resM.Data, expectedLength)
}
func TestApiGetRandomWordsFailed(t *testing.T) {
	s, mock := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()

	query := "SELECT * FROM word ORDER BY RANDOM() LIMIT 3"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(sql.ErrConnDone)

	server := NewApiServer(os.Getenv("API_HOST")+":"+os.Getenv("API_PORT"), s)
	router := server.SetupRouter()
	router = GetRandomWords(server, router)
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	req, _ := http.NewRequest("GET", "/random-words", nil)
	req.Header.Add("Accept", "application/json")
	router.ServeHTTP(w, req)
	var resM ResponseModel[[]*Word]
	decoder := json.NewDecoder(w.Body)

	err := decoder.Decode(&resM)
	if err != nil {
		t.Errorf("Json decoder doesn't work: %s", err)
	}

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, sql.ErrConnDone.Error(), resM.Error)
}
func TestApiGetWord(t *testing.T) {
	s, mock := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()
	rows := sqlmock.NewRows([]string{"id", "created_at", "main_language", "foreign_language"}).AddRow(wordModel.Id, wordModel.Created_at, wordModel.Main_language, wordModel.Foreign_language)
	query := "SELECT * FROM word WHERE id = $1"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(wordModel.Id).WillReturnRows(rows)

	server := NewApiServer(os.Getenv("API_HOST")+":"+os.Getenv("API_PORT"), s)
	router := server.SetupRouter()
	router = GetWord(server, router)
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	req, _ := http.NewRequest("GET", "/words/"+strconv.Itoa(wordModel.Id), nil)
	req.Header.Add("Accept", "application/json")
	router.ServeHTTP(w, req)
	var resM ResponseModel[*Word]
	decoder := json.NewDecoder(w.Body)

	err := decoder.Decode(&resM)
	if err != nil {
		t.Errorf("Json decoder doesn't work: %s", err)
	}
	if resM.Data == nil {
		t.Errorf("Response is nil")
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, wordModel.Id, resM.Data.Id)
}
func TestApiGetWord_InvalidIdError(t *testing.T) {
	s, _ := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()

	server := NewApiServer(os.Getenv("API_HOST")+":"+os.Getenv("API_PORT"), s)
	router := server.SetupRouter()
	router = GetWord(server, router)
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	req, _ := http.NewRequest("GET", "/words/"+"test", nil)
	req.Header.Add("Accept", "application/json")
	router.ServeHTTP(w, req)
	var resM ResponseModel[*Word]
	decoder := json.NewDecoder(w.Body)

	err := decoder.Decode(&resM)
	if err != nil {
		t.Errorf("Json decoder doesn't work: %s", err)
	}

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "invalid id is given test", resM.Error)
}
func TestApiGetWordFailed_DB(t *testing.T) {
	s, mock := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()
	query := "SELECT * FROM word WHERE id = $1"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(wordModel.Id).WillReturnError(sql.ErrConnDone)

	server := NewApiServer(os.Getenv("API_HOST")+":"+os.Getenv("API_PORT"), s)
	router := server.SetupRouter()
	router = GetWord(server, router)
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	req, _ := http.NewRequest("GET", "/words/"+strconv.Itoa(wordModel.Id), nil)
	req.Header.Add("Accept", "application/json")
	router.ServeHTTP(w, req)
	var resM ResponseModel[*Word]
	decoder := json.NewDecoder(w.Body)

	err := decoder.Decode(&resM)
	if err != nil {
		t.Errorf("Json decoder doesn't work: %s", err)
	}

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, sql.ErrConnDone.Error(), resM.Error)
}
func TestApiPostWord(t *testing.T) {
	s, mock := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()
	rows := sqlmock.NewRows([]string{"id"}).AddRow(wordModel.Id)
	query := "INSERT INTO word (main_language, foreign_language, created_at) VALUES ($1, $2, $3) RETURNING id"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(wordModel.Main_language, wordModel.Foreign_language, AnyTime{}).WillReturnRows(rows)

	server := NewApiServer(os.Getenv("API_HOST")+":"+os.Getenv("API_PORT"), s)
	router := server.SetupRouter()
	router = PostWord(server, router)
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")

	req, _ := http.NewRequest("POST", "/words", strings.NewReader(string(wordJson)))
	req.Header.Add("Accept", "application/json")
	router.ServeHTTP(w, req)
	var resM ResponseModel[*Word]
	decoder := json.NewDecoder(w.Body)

	err := decoder.Decode(&resM)
	if err != nil {
		t.Errorf("Json decoder doesn't work: %s", err)
	}
	if resM.Data == nil {
		t.Errorf("Response is nil")
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, wordModel.Id, resM.Data.Id)
	assert.Equal(t, wordModel.Main_language, resM.Data.Main_language)
	assert.Equal(t, wordModel.Foreign_language, resM.Data.Foreign_language)
}
func TestApiPostWordFailed_NilControl(t *testing.T) {
	m1 := Word{
		Main_language:    "",
		Foreign_language: "text",
	}
	m2 := Word{
		Main_language:    "text",
		Foreign_language: "",
	}
	testCases := []struct {
		name        string
		model       Word
		expectedErr error
	}{
		{
			name:        "TestApiPostWordMainLanguageIsEmpty",
			model:       m1,
			expectedErr: fmt.Errorf("main language is empty"),
		},
		{
			name:        "TestApiPostWordForeignLanguageIsEmpty",
			model:       m2,
			expectedErr: fmt.Errorf("foreign language is empty"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			wordJ, _ := json.Marshal(tc.model)

			s, _ := MockCreatePostgreSqlConnection()
			defer func() {
				s.db.Close()
			}()
			server := NewApiServer(os.Getenv("API_HOST")+":"+os.Getenv("API_PORT"), s)
			router := server.SetupRouter()
			router = PostWord(server, router)
			w := httptest.NewRecorder()
			w.Header().Set("Content-Type", "application/json")

			req, _ := http.NewRequest("POST", "/words", strings.NewReader(string(wordJ)))
			req.Header.Add("Accept", "application/json")
			router.ServeHTTP(w, req)
			var resM ResponseModel[*Word]
			decoder := json.NewDecoder(w.Body)

			err := decoder.Decode(&resM)
			if err != nil {
				t.Errorf("Json decoder doesn't work: %s", err)
			}
			assert.Equal(t, 400, w.Code)
			assert.Equal(t, tc.expectedErr.Error(), resM.Error)
		})
	}
}
func TestApiPostWordFailed_DB(t *testing.T) {
	s, mock := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()
	query := "INSERT INTO word (main_language, foreign_language, created_at) VALUES ($1, $2, $3) RETURNING id"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(wordModel.Main_language, wordModel.Foreign_language, AnyTime{}).WillReturnError(sql.ErrConnDone)

	server := NewApiServer(os.Getenv("API_HOST")+":"+os.Getenv("API_PORT"), s)
	router := server.SetupRouter()
	router = PostWord(server, router)
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")

	req, _ := http.NewRequest("POST", "/words", strings.NewReader(string(wordJson)))
	req.Header.Add("Accept", "application/json")
	router.ServeHTTP(w, req)
	var resM ResponseModel[*Word]
	decoder := json.NewDecoder(w.Body)

	err := decoder.Decode(&resM)
	if err != nil {
		t.Errorf("Json decoder doesn't work: %s", err)
	}

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, sql.ErrConnDone.Error(), resM.Error)
}
func TestApiUpdateWord(t *testing.T) {
	newWord := Word{
		Id:               wordModel.Id,
		Main_language:    wordModel.Main_language,
		Foreign_language: wordModel.Foreign_language,
	}
	wordJ, _ := json.Marshal(newWord)

	s, mock := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()

	row := sqlmock.NewRows([]string{"exists"}).AddRow(true)
	query1 := "SELECT exists (SELECT 1 FROM word WHERE id=$1 LIMIT 1)"
	mock.ExpectQuery(regexp.QuoteMeta(query1)).WithArgs(newWord.Id).WillReturnRows(row)

	rows := sqlmock.NewRows([]string{"id", "created_at", "main_language", "foreign_language"})
	query2 := "UPDATE word SET main_language=$1, foreign_language=$2 WHERE id=$3"
	mock.ExpectQuery(regexp.QuoteMeta(query2)).WithArgs(newWord.Main_language, newWord.Foreign_language, newWord.Id).WillReturnRows(rows)

	server := NewApiServer(os.Getenv("API_HOST")+":"+os.Getenv("API_PORT"), s)
	router := server.SetupRouter()
	router = UpdateWord(server, router)
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")

	req, _ := http.NewRequest("PUT", "/words", strings.NewReader(string(wordJ)))
	req.Header.Add("Accept", "application/json")

	router.ServeHTTP(w, req)
	var resM ResponseModel[*Word]
	decoder := json.NewDecoder(w.Body)

	err := decoder.Decode(&resM)
	if err != nil {
		t.Errorf("Json decoder doesn't work: %s", err)
	}

	assert.Equal(t, 200, w.Code)
}
func TestApiUpdateWordFailed_NilControl(t *testing.T) {
	m1 := Word{
		Id:               wordModel.Id,
		Main_language:    "",
		Foreign_language: "text",
	}
	m2 := Word{
		Id:               wordModel.Id,
		Main_language:    "text",
		Foreign_language: "",
	}
	testCases := []struct {
		name        string
		model       Word
		expectedErr error
	}{
		{
			name:        "TestApiPostWordMainLanguageIsEmpty",
			model:       m1,
			expectedErr: fmt.Errorf("main language is empty"),
		},
		{
			name:        "TestApiPostWordForeignLanguageIsEmpty",
			model:       m2,
			expectedErr: fmt.Errorf("foreign language is empty"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			wordJ, _ := json.Marshal(tc.model)

			s, _ := MockCreatePostgreSqlConnection()
			defer func() {
				s.db.Close()
			}()

			server := NewApiServer(os.Getenv("API_HOST")+":"+os.Getenv("API_PORT"), s)
			router := server.SetupRouter()
			router = UpdateWord(server, router)
			w := httptest.NewRecorder()
			w.Header().Set("Content-Type", "application/json")

			req, _ := http.NewRequest("PUT", "/words", strings.NewReader(string(wordJ)))
			req.Header.Add("Accept", "application/json")

			router.ServeHTTP(w, req)
			var resM ResponseModel[*Word]
			decoder := json.NewDecoder(w.Body)

			err := decoder.Decode(&resM)
			if err != nil {
				t.Errorf("Json decoder doesn't work: %s", err)
			}

			assert.Equal(t, 400, w.Code)
			assert.Equal(t, tc.expectedErr.Error(), resM.Error)
		})
	}
}
func TestApiUpdateWordFailed_DBExists(t *testing.T) {
	newWord := Word{
		Id:               wordModel.Id,
		Main_language:    wordModel.Main_language,
		Foreign_language: wordModel.Foreign_language,
	}
	wordJ, _ := json.Marshal(newWord)

	s, mock := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()

	query1 := "SELECT exists (SELECT 1 FROM word WHERE id=$1 LIMIT 1)"
	mock.ExpectQuery(regexp.QuoteMeta(query1)).WithArgs(newWord.Id).WillReturnError(sql.ErrConnDone)

	server := NewApiServer(os.Getenv("API_HOST")+":"+os.Getenv("API_PORT"), s)
	router := server.SetupRouter()
	router = UpdateWord(server, router)
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")

	req, _ := http.NewRequest("PUT", "/words", strings.NewReader(string(wordJ)))
	req.Header.Add("Accept", "application/json")

	router.ServeHTTP(w, req)
	var resM ResponseModel[*Word]
	decoder := json.NewDecoder(w.Body)

	err := decoder.Decode(&resM)
	if err != nil {
		t.Errorf("Json decoder doesn't work: %s", err)
	}
	assert.Equal(t, 404, w.Code)
	assert.Equal(t, sql.ErrConnDone.Error(), resM.Error)
}
func TestApiUpdateWordFailed_DBUpdate(t *testing.T) {
	newWord := Word{
		Id:               wordModel.Id,
		Main_language:    wordModel.Main_language,
		Foreign_language: wordModel.Foreign_language,
	}
	wordJ, _ := json.Marshal(newWord)

	s, mock := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()

	row := sqlmock.NewRows([]string{"exists"}).AddRow(true)
	query1 := "SELECT exists (SELECT 1 FROM word WHERE id=$1 LIMIT 1)"
	mock.ExpectQuery(regexp.QuoteMeta(query1)).WithArgs(newWord.Id).WillReturnRows(row)

	query2 := "UPDATE word SET main_language=$1, foreign_language=$2 WHERE id=$3"
	mock.ExpectQuery(regexp.QuoteMeta(query2)).WithArgs(newWord.Main_language, newWord.Foreign_language, newWord.Id).WillReturnError(sql.ErrConnDone)

	server := NewApiServer(os.Getenv("API_HOST")+":"+os.Getenv("API_PORT"), s)
	router := server.SetupRouter()
	router = UpdateWord(server, router)
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")

	req, _ := http.NewRequest("PUT", "/words", strings.NewReader(string(wordJ)))
	req.Header.Add("Accept", "application/json")

	router.ServeHTTP(w, req)
	var resM ResponseModel[*Word]
	decoder := json.NewDecoder(w.Body)

	err := decoder.Decode(&resM)
	if err != nil {
		t.Errorf("Json decoder doesn't work: %s", err)
	}
	assert.Equal(t, 404, w.Code)
	assert.Equal(t, sql.ErrConnDone.Error(), resM.Error)
}
func TestApiDeleteWord(t *testing.T) {
	s, mock := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()

	query := "DELETE FROM word WHERE id=$1"
	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(wordModel.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	server := NewApiServer(os.Getenv("API_HOST")+":"+os.Getenv("API_PORT"), s)
	router := server.SetupRouter()
	router = DeleteWord(server, router)
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")

	req, _ := http.NewRequest("DELETE", "/words/"+strconv.Itoa(wordModel.Id), nil)
	req.Header.Add("Accept", "application/json")

	router.ServeHTTP(w, req)
	var resM ResponseModel[*Word]
	decoder := json.NewDecoder(w.Body)

	err := decoder.Decode(&resM)
	if err != nil {
		t.Errorf("Json decoder doesn't work: %s", err)
	}

	assert.Equal(t, 200, w.Code)
}
func TestApiDeleteWordFailed_GetId(t *testing.T) {
	s, _ := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()

	server := NewApiServer(os.Getenv("API_HOST")+":"+os.Getenv("API_PORT"), s)
	router := server.SetupRouter()
	router = DeleteWord(server, router)
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")

	req, _ := http.NewRequest("DELETE", "/words/test", nil)
	req.Header.Add("Accept", "application/json")

	router.ServeHTTP(w, req)
	var resM ResponseModel[*Word]
	decoder := json.NewDecoder(w.Body)

	err := decoder.Decode(&resM)
	if err != nil {
		t.Errorf("Json decoder doesn't work: %s", err)
	}

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, "invalid id is given test", resM.Error)
}
func TestApiDeleteWordFailed_DB(t *testing.T) {
	s, mock := MockCreatePostgreSqlConnection()
	defer func() {
		s.db.Close()
	}()

	query := "DELETE FROM word WHERE id=$1"
	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(wordModel.Id).WillReturnResult(sqlmock.NewResult(0, 0))

	server := NewApiServer(os.Getenv("API_HOST")+":"+os.Getenv("API_PORT"), s)
	router := server.SetupRouter()
	router = DeleteWord(server, router)
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")

	req, _ := http.NewRequest("DELETE", "/words/"+strconv.Itoa(wordModel.Id), nil)
	req.Header.Add("Accept", "application/json")

	router.ServeHTTP(w, req)
	var resM ResponseModel[*Word]
	decoder := json.NewDecoder(w.Body)

	err := decoder.Decode(&resM)
	if err != nil {
		t.Errorf("Json decoder doesn't work: %s", err)
	}
	errorMsg := "id " + strconv.Itoa(wordModel.Id) + " does not exist"
	assert.Equal(t, 404, w.Code)
	assert.Equal(t, errorMsg, resM.Error)
}
