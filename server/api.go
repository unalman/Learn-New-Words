package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ResponseModel[T interface{ []*Word | *Word | string }] struct {
	Data  T      `json:"data"`
	Error string `json:"error"`
}

type APIServer struct {
	listenAddr string
	store      Storage
}

// Return APIServer model
func NewApiServer(listenAddr string, storage Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      storage,
	}
}

// API routers
func (s *APIServer) Run() {
	router := gin.Default()
	router.GET("/words", s.GetWords)
	router.GET("/words/:id", s.GetWord)
	router.POST("/words", s.PostWord)
	router.PUT("/words", s.UpdateWord)
	router.DELETE("/words/:id", s.DeleteWord)

	router.Run(s.listenAddr)
}

// Get Word List model
func (s *APIServer) GetWords(c *gin.Context) {
	var responseModel ResponseModel[[]*Word]
	word, err := s.store.GetListWord()
	if err != nil {
		responseModel.Error = err.Error()
		c.IndentedJSON(http.StatusNotFound, responseModel)
		return
	}
	responseModel.Data = word
	c.IndentedJSON(http.StatusOK, responseModel)
}

// Get Word model
func (s *APIServer) GetWord(c *gin.Context) {
	var responseModel ResponseModel[*Word]
	id, err := getId(c)
	if err != nil {
		responseModel.Error = err.Error()
		c.IndentedJSON(http.StatusBadRequest, responseModel)
		return
	}
	word, err := s.store.GetWordById(id)
	if err != nil {
		responseModel.Error = err.Error()
		c.IndentedJSON(http.StatusNotFound, responseModel)
		return
	}
	responseModel.Data = word
	c.IndentedJSON(http.StatusOK, responseModel)
}

// Insert Word model
func (s *APIServer) PostWord(c *gin.Context) {
	var responseModel ResponseModel[*Word]

	word, err := checkWordTextValidation(c)
	if err != nil {
		responseModel.Error = err.Error()
		c.IndentedJSON(http.StatusBadRequest, responseModel)
		return
	}

	newWord, err := s.store.CreateWord(word.Main_language, word.Foreign_language)
	if err != nil {
		responseModel.Error = err.Error()
		c.IndentedJSON(http.StatusNotFound, responseModel)
		return
	}
	responseModel.Data = newWord
	c.IndentedJSON(http.StatusOK, responseModel)
}

// Update a Word model
func (s *APIServer) UpdateWord(c *gin.Context) {
	var responseModel ResponseModel[*Word]

	word, err := checkWordTextValidation(c)
	if err != nil {
		responseModel.Error = err.Error()
		c.IndentedJSON(http.StatusBadRequest, responseModel)
		return
	}
	err = s.checkIdExistance(word)
	if err != nil {
		responseModel.Error = err.Error()
		c.IndentedJSON(http.StatusNotFound, responseModel)
		return
	}
	err = s.store.UpdateWord(word)
	if err != nil {
		responseModel.Error = err.Error()
		c.IndentedJSON(http.StatusNotFound, responseModel)
		return
	}
	c.IndentedJSON(http.StatusOK, responseModel)
}

// Delete a row from Word table
func (s *APIServer) DeleteWord(c *gin.Context) {
	var responseModel ResponseModel[*Word]
	id, err := getId(c)
	if err != nil {
		responseModel.Error = err.Error()
		c.IndentedJSON(http.StatusBadRequest, responseModel)
		return
	}
	err = s.store.DeleteWord(id)
	if err != nil {
		responseModel.Error = err.Error()
		c.IndentedJSON(http.StatusNotFound, responseModel)
		return
	}
	c.IndentedJSON(http.StatusOK, responseModel)
}

// Parsing id string to int
func getId(c *gin.Context) (int, error) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return id, fmt.Errorf("invalid id is given %s", idString)
	}
	return id, nil
}

// Word model fields validation function
func checkWordTextValidation(c *gin.Context) (*Word, error) {
	var word *Word
	c.BindJSON(&word)
	if checkWordValidation(word.Main_language) {
		return nil, fmt.Errorf("main language is null")
	}
	if checkWordValidation(word.Foreign_language) {
		return nil, fmt.Errorf("foreign language is null")
	}
	return word, nil
}

// string validation funtion
func checkWordValidation(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}

// Checking id existance at the word table
func (s *APIServer) checkIdExistance(word *Word) error {
	isExists, err := s.store.ExistsWord(word.Id)
	if err != nil {
		return err
	}
	if !isExists {
		return fmt.Errorf("id %s does not exists", strconv.Itoa(word.Id))
	}
	return nil
}
