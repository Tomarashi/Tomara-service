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

	queryParamSubWord   = "sub_word"
	queryParamWordN     = "word_n"
	queryParamRequestId = "request_id"

	allowOriginHeaderName = "Access-Control-Allow-Origin"
)

type WordController struct {
	Repository repository.ITomaraRepository
}

type GetWordsResponse struct {
	Words      []string `json:"words"`
	TakenNanos int64    `json:"taken_ns"`
}

type GetWordsResponseWithReqId struct {
	Words      []string `json:"words"`
	TakenNanos int64    `json:"taken_ns"`
	RequestId  string   `json:"request_id"`
}

func (w WordController) Greetings(c *gin.Context) {
	c.Header(allowOriginHeaderName, "*")
	c.String(http.StatusOK, "Hello!")
}

func (w WordController) GetWords(c *gin.Context) {
	subWord, exists := c.GetQuery(queryParamSubWord)
	c.Header(allowOriginHeaderName, "*")
	if !exists {
		c.String(http.StatusBadRequest, HttpParameterNotExist(queryParamSubWord))
		return
	}
	if subWord == "" {
		c.JSON(http.StatusOK, GetWordsResponse{Words: []string{}, TakenNanos: 0})
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
	if requestId, exists := c.GetQuery(queryParamRequestId); exists {
		c.JSON(http.StatusOK, GetWordsResponseWithReqId{Words: result, TakenNanos: takenTime, RequestId: requestId})
	} else {
		c.JSON(http.StatusOK, GetWordsResponse{Words: result, TakenNanos: takenTime})
	}
}
