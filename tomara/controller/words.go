package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tomara-service/tomara/repository"
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

type GreetingsResponse struct {
	Msg    string `json:"message"`
	DBType string `json:"db_type"`
}

type GetWordsFailResponse struct {
	ErrorMsg string `json:"error_msg"`
}

type GetWordsResponse struct {
	Words []string `json:"words"`
}

type GetWordsResponseWithReqId struct {
	Words     []string `json:"words"`
	RequestId string   `json:"request_id"`
}

func (w WordController) Greetings(c *gin.Context) {
	c.Header(allowOriginHeaderName, "*")
	c.JSON(http.StatusOK, GreetingsResponse{
		Msg:    "Hello!",
		DBType: w.Repository.DBType(),
	})
}

func (w WordController) GetWords(c *gin.Context) {
	subWord, exists := c.GetQuery(queryParamSubWord)
	c.Header(allowOriginHeaderName, "*")
	if !exists {
		c.String(http.StatusBadRequest, HttpParameterNotExist(queryParamSubWord))
		return
	}
	if subWord == "" {
		c.JSON(http.StatusOK, GetWordsResponse{Words: []string{}})
		return
	}
	wordNumber := defaultWordN
	if wordNArgStr, ok := c.GetQuery(queryParamWordN); ok {
		if wordNArg, err := strconv.Atoi(wordNArgStr); err == nil {
			wordNumber = wordNArg
		}
	}
	err, result := w.Repository.GetWordsStartsWith(subWord, wordNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, GetWordsFailResponse{ErrorMsg: err.Error()})
		return
	}
	if requestId, exists := c.GetQuery(queryParamRequestId); exists {
		c.JSON(http.StatusOK, GetWordsResponseWithReqId{Words: result, RequestId: requestId})
	} else {
		c.JSON(http.StatusOK, GetWordsResponse{Words: result})
	}
}
