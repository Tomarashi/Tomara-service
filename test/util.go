package test

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func get(url string) *http.Request {
	return getWithBody(url, nil)
}

func getWithBody(url string, body map[string]interface{}) *http.Request {
	bodyBuffer := new(bytes.Buffer)
	err := json.NewEncoder(bodyBuffer).Encode(body)
	if err != nil {
		return nil
	}

	request, err := http.NewRequest("GET", url, bodyBuffer)
	if err != nil {
		panic(err)
	}
	return request
}

func contains[T string](elements []T, elem T) bool {
	for _, e := range elements {
		if e == elem {
			return true
		}
	}
	return false
}
