package tomara

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tomara-service/tomara/server"
)

func Start(host string, port int) {
	gin.SetMode(gin.ReleaseMode)

	app := server.GetServerDefault()
	err := app.Run(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		panic(err)
	}
}
