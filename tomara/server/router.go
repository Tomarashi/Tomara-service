package server

import (
	"github.com/gin-gonic/gin"
	"tomara-service/tomara/repository"
)

const (
	ApiRootPath = "/extension/api"
)

func GetServerDefault() *gin.Engine {
	tomaraRepository := repository.MakeMySqlRepositoryDefaultConfig()
	return GetServer(tomaraRepository)
}

func GetServer(tomaraRepository repository.ITomaraRepository) *gin.Engine {
	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	router := server.Group(ApiRootPath)
	SetUpWordsRouter(router, tomaraRepository)

	PrintRoutes(server)

	return server
}
