package server

import (
	"github.com/gin-gonic/gin"
	"tomara-service/tomara/controller"
	"tomara-service/tomara/repository"
)

func SetUpWordsRouter(mainRoute *gin.RouterGroup, tomaraRepository repository.ITomaraRepository) {
	wordsController := new(controller.WordController)
	wordsController.Repository = tomaraRepository

	wordsRouter := mainRoute.Group("/word")
	{
		wordsRouter.GET("/greetings", wordsController.Greetings)
		wordsRouter.GET("/get", wordsController.GetWords)
	}
}
