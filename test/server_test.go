package test

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"math"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"tomara-service/tomara/repository"
	"tomara-service/tomara/server"
)

var tomaraServer *gin.Engine

type TestResponse struct {
	Words []string `json:"words"`
	Time  int      `json:"taken_ns"`
}

func (r *TestResponse) Parse(value []byte) {
	err := json.Unmarshal(value, r)
	if err != nil {
		panic(err)
	}
}

func init() {
	gin.SetMode(gin.ReleaseMode)

	fakeRepository := repository.MakeFakeRepositoryFromFile(TestDataFilePath)
	tomaraServer = server.GetServer(fakeRepository)
}

func TestGreetings(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := get("/extension/api/word/greetings")

	tomaraServer.ServeHTTP(recorder, request)
	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "Hello!", recorder.Body.String())
}

func TestGet(t *testing.T) {
	var testResponse TestResponse
	var recorder *httptest.ResponseRecorder
	var request *http.Request

	recorder = httptest.NewRecorder()
	request = get(formatUri(map[string]interface{}{
		"sub_word": "ა",
		"word_n":   math.MaxInt,
	}))

	tomaraServer.ServeHTTP(recorder, request)
	testResponse.Parse(recorder.Body.Bytes())
	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, 213, len(testResponse.Words))

	recorder = httptest.NewRecorder()
	request = get(formatUri(map[string]interface{}{
		"sub_word": "ბი",
		"word_n":   math.MaxInt,
	}))

	tomaraServer.ServeHTTP(recorder, request)
	testResponse.Parse(recorder.Body.Bytes())
	assert.Equal(t, 200, recorder.Code)
	assert.True(t, len(testResponse.Words) > 0)
	for _, word := range testResponse.Words {
		ind := strings.Index(word, "ბი")
		assert.Equal(t, 0, ind)
	}
}

func formatUri(params map[string]interface{}) string {
	var builder strings.Builder
	builder.WriteString("/extension/api/word/get?")
	addAmpersand := false
	for key, value := range params {
		if addAmpersand {
			builder.WriteRune('&')
		}
		builder.WriteString(key)
		builder.WriteRune('=')

		var writableVal string
		switch value.(type) {
		case int:
			writableVal = strconv.Itoa(value.(int))
		case string:
			writableVal = url.QueryEscape(value.(string))
		default:
			panic(fmt.Errorf("invalid type"))
		}
		builder.WriteString(writableVal)

		addAmpersand = true
	}
	return builder.String()
}
