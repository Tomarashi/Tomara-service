package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	defaultWordN = 8

	queryParamSubWord = "sub_word"
	queryParamWordN   = "word_n"
)

type WordController struct{}

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
	c.String(http.StatusOK, "%s, %d", subWord, wordNumber)
}
