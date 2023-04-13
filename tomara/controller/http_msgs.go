package controller

import (
	"fmt"
)

func HttpParameterNotExist(paramName string) string {
	return fmt.Sprintf("Parameter '%s' doesn't exist", paramName)
}
