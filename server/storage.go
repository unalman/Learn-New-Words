package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateWord(mainlanguage, foreignLanguage string) (*Word, error)
	GetWordById(id int) (*Word, error)
	GetListWord() ([]*Word, error)
	UpdateWord(word *Word) error
	DeleteWord(id int) error
	ExistsWord(id int) (bool, error)
}

type PostgreSqlInfo struct {
	db *sql.DB
}
type Word struct {
	Id               int       `json:"id"`
	Created_at       time.Time `json:"created_at"`
	Main_language    string    `json:"main_language"`
	Foreign_language string    `json:"foreign_language"`
}

// PostgreSql connect
func CreatePostgreSqlConnection() (*PostgreSqlInfo, error) {
	postgresqlDbInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", postgresqlDbInfo)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgreSqlInfo{
		db: db,
	}, nil
}

// Inserting the Word model to the Word table
func (s *PostgreSqlInfo) CreateWord(mainlanguage, foreignLanguage string) (*Word, error) {
	word := NewWordMapper(mainlanguage, foreignLanguage)
	rows, err := s.db.Query("INSERT INTO word (main_language, foreign_language, created_at) VALUES ($1, $2, $3) RETURNING id",
		word.Main_language, word.Foreign_language, word.Created_at)
	if err != nil {
		return nil, err
	}
	var idString string
	for rows.Next() {
		rows.Scan(&idString)
	}
	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, err
	}
	word.Id = id
	return word, err
}

// Get a row from Word table
func (s *PostgreSqlInfo) GetWordById(id int) (*Word, error) {
	rows, err := s.db.Query("SELECT * FROM word WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoWord(rows)
	}
	defer rows.Close()

	return nil, fmt.Errorf("id %s does not exist", strconv.Itoa(id))
}

// Get row list from Word table
func (s *PostgreSqlInfo) GetListWord() ([]*Word, error) {
	rows, err := s.db.Query("SELECT * FROM word")
	if err != nil {
		return nil, err
	}
	words := []*Word{}
	for rows.Next() {
		newWord, err := scanIntoWord(rows)
		if err != nil {
			return nil, err
		}
		words = append(words, newWord)
	}
	defer rows.Close()
	return words, err
}

// Update a row from Word table
func (s *PostgreSqlInfo) UpdateWord(word *Word) error {
	_, err := s.db.Query("UPDATE word SET main_language=$1, foreign_language=$2 WHERE id=$3", word.Main_language, word.Foreign_language, word.Id)
	if err != nil {
		return err
	}
	return nil
}

// Deleting a row from Word table
func (s *PostgreSqlInfo) DeleteWord(id int) error {
	res, err := s.db.Exec("DELETE FROM word WHERE id=$1", id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("id %s does not exist", strconv.Itoa(id))
	}
	return err
}

// Existance control in the Word table
func (s *PostgreSqlInfo) ExistsWord(id int) (bool, error) {
	var isExists bool
	rows, err := s.db.Query("SELECT exists (SELECT 1 FROM word WHERE id=$1 LIMIT 1)", id)
	if err != nil {
		return isExists, err
	}

	for rows.Next() {
		rows.Scan(&isExists)
	}
	defer rows.Close()
	return isExists, nil
}

// Word model mapping with SQL response
func scanIntoWord(rows *sql.Rows) (*Word, error) {
	word := new(Word)
	err := rows.Scan(
		&word.Id,
		&word.Created_at,
		&word.Main_language,
		&word.Foreign_language)
	return word, err
}

// Word model is mapping with parameters
func NewWordMapper(mainlanguage, foreignLanguage string) *Word {
	return &Word{
		Created_at:       time.Now().UTC(),
		Main_language:    mainlanguage,
		Foreign_language: foreignLanguage,
	}
}
