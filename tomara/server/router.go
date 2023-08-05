package server

import (
	"github.com/gin-gonic/gin"
	"tomara-service/tomara/repository"
)

const (
	ApiRootPath = "/extension/api"
)

func GetServer() *gin.Engine {
	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	tomaraRepository := repository.MakeMySqlRepositoryDefaultConfig()

	router := server.Group(ApiRootPath)
	SetUpWordsRouter(router, tomaraRepository)

	PrintRoutes(server)

	return server
}
