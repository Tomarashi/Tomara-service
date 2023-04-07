package server

import (
	"github.com/gin-gonic/gin"
	"tomara-service/tomara/controller"
)

func SetUpWordsRouter(mainRoute *gin.RouterGroup) {
	wordsController := new(controller.WordController)

	wordsRouter := mainRoute.Group("/word")
    {
        wordsRouter.GET("/", wordsController.Hello)
    }
}
