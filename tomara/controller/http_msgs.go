package controller

import (
	"fmt"
)

type ErrorMessage struct {
	Msg    string `json:"message"`
}

func HttpParameterNotExist(paramName string) ErrorMessage {
	msg := fmt.Sprintf("Parameter '%s' doesn't exist", paramName)
	return ErrorMessage{Msg: msg}
}
