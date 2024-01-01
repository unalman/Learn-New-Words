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

func NewApiServer(listenAddr string, storage Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      storage,
	}
}
func (s *APIServer) Run() {
	router := gin.Default()
	router.GET("/words", s.GetWords)
	router.GET("/words/:id", s.GetWord)
	router.POST("/words", s.PostWord)
	router.PUT("/words", s.UpdateWord)
	router.DELETE("/words/:id", s.DeleteWord)

	router.Run(s.listenAddr)
}

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
func (s *APIServer) UpdateWord(c *gin.Context) {
	var responseModel ResponseModel[*Word]

	word, err := checkWordTextValidation(c)
	if err != nil {
		responseModel.Error = err.Error()
		c.IndentedJSON(http.StatusBadRequest, responseModel)
		return
	}
	err = s.checkIdIsExist(word)
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

func checkWordValidation(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}
func getId(c *gin.Context) (int, error) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return id, fmt.Errorf("invalid id is given %s", idString)
	}
	return id, nil
}
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
func (s *APIServer) checkIdIsExist(word *Word) error {
	isExists, err := s.store.ExistsWord(word.Id)
	if err != nil {
		return err
	}
	if !isExists {
		return fmt.Errorf("id %s does not exists", strconv.Itoa(word.Id))
	}
	return nil
}
