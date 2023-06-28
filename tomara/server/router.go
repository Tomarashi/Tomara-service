package server

import (
	"github.com/gin-gonic/gin"
	"tomara-service/tomara/repository"
)

func GetServer() *gin.Engine {
	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	tomaraRepository := repository.MakeFakeRepositoryFromFile("data/dev-words.txt")

	router := server.Group("/extension/api")
	SetUpWordsRouter(router, tomaraRepository)

	return server
}
