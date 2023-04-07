package server

import (
	"github.com/gin-gonic/gin"
)

func GetServer() *gin.Engine {
    server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	router := server.Group("/")
	SetUpWordsRouter(router)

	return server
}
