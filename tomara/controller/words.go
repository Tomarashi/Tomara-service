package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tomara-service/tomara/repository"
	"tomara-service/tomara/utils"
)

const (
	defaultWordN = 8

	queryParamSubWord = "sub_word"
	queryParamWordN   = "word_n"
)

type WordController struct {
	Repository repository.ITomaraRepository
}

type GetWordsResponse struct {
	Words      []string `json:"words"`
	TakenNanos int64    `json:"taken_ns"`
}

func (w WordController) GetWords(c *gin.Context) {
	subWord, exists := c.GetQuery(queryParamSubWord)
	if !exists {
		c.String(http.StatusBadRequest, HttpParameterNotExist(queryParamSubWord))
		return
	}
	wordNumber := defaultWordN
	if wordNArgStr, ok := c.GetQuery(queryParamWordN); ok {
		if wordNArg, err := strconv.Atoi(wordNArgStr); err == nil {
			wordNumber = wordNArg
		}
	}
	startTime := utils.CurrentNanos()
	result := w.Repository.GetWordsStartsWith(subWord, wordNumber)
	takenTime := utils.FromTimeInNanos(startTime)
	c.JSON(http.StatusOK, GetWordsResponse{Words: result, TakenNanos: takenTime})
}
