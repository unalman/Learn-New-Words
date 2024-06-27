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
func (s *APIServer) SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// Get Word List model
func GetWords(s *APIServer, r *gin.Engine) *gin.Engine {
	r.GET("/words", func(c *gin.Context) {
		var responseModel ResponseModel[[]*Word]
		word, err := s.store.GetListWord()
		if err != nil {
			responseModel.Error = err.Error()
			c.IndentedJSON(http.StatusNotFound, responseModel)
			return
		}
		responseModel.Data = word
		c.IndentedJSON(http.StatusOK, responseModel)
	})
	return r
}

// Get Word model
func GetWord(s *APIServer, r *gin.Engine) *gin.Engine {
	r.GET("/words/:id", func(c *gin.Context) {
		var responseModel ResponseModel[*Word]
		idString := c.Param("id")
		id, err := getId(idString)
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
	})
	return r
}

// Insert Word model
func PostWord(s *APIServer, r *gin.Engine) *gin.Engine {
	r.POST("/words", func(c *gin.Context) {
		var responseModel ResponseModel[*Word]
		var word *Word
		c.BindJSON(&word)
		err := checkWordTextValidation(word)
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
	})
	return r
}

// Update a Word model
func UpdateWord(s *APIServer, r *gin.Engine) *gin.Engine {
	r.PUT("/words", func(c *gin.Context) {
		var responseModel ResponseModel[*Word]
		var word *Word
		c.BindJSON(&word)
		err := checkWordTextValidation(word)
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
	})
	return r
}

// Delete a row from Word table
func DeleteWord(s *APIServer, r *gin.Engine) *gin.Engine {
	r.DELETE("/words/:id", func(c *gin.Context) {
		var responseModel ResponseModel[*Word]
		idString := c.Param("id")
		id, err := getId(idString)
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
	})
	return r
}

// Parsing id string to int
func getId(idString string) (int, error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return id, fmt.Errorf("invalid id is given %s", idString)
	}
	return id, nil
}

// Word model fields validation function
func checkWordTextValidation(word *Word) error {
	if checkWordValidation(word.Main_language) {
		return fmt.Errorf("main language is empty")
	}
	if checkWordValidation(word.Foreign_language) {
		return fmt.Errorf("foreign language is empty")
	}
	return nil
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
