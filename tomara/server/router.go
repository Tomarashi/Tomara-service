package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tomara-service/tomara/repository"
)

const (
	DoPrintRoutes = false

	ApiRootPath = "/extension/api"
)

func GetServerDefault() *gin.Engine {
	tomaraRepository := repository.GetFirstValidRepository()
	fmt.Printf("Repository type: %s\n", tomaraRepository.DBType())
	return GetServer(tomaraRepository)
}

func GetServer(tomaraRepository repository.ITomaraRepository) *gin.Engine {
	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	router := server.Group(ApiRootPath)
	SetUpWordsRouter(router, tomaraRepository)

	if DoPrintRoutes {
		PrintRoutes(server)
	}

	return server
}
