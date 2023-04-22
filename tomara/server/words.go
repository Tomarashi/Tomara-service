package server

import (
	"github.com/gin-gonic/gin"
	"tomara-service/tomara/controller"
	"tomara-service/tomara/repository"
)

func SetUpWordsRouter(mainRoute *gin.RouterGroup) {
	wordsController := new(controller.WordController)
	wordsController.Repository = repository.MakeFakeRepositoryFromFile("data/dev-words.txt")

	wordsRouter := mainRoute.Group("/word")
	{
		wordsRouter.GET("/get", wordsController.GetWords)
	}
}
