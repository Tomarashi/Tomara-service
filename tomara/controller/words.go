package controller

import (
	"github.com/gin-gonic/gin"
)

type WordController struct {}

func (w WordController) Hello(c *gin.Context) {
	c.String(200, "Hello World!")
}
