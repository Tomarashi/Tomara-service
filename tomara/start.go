package tomara

import (
	"fmt"
	"tomara-service/tomara/server"
)

func Start(host string, port int) {
	app := server.GetServerDefault()
	err := app.Run(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		panic(err)
	}
}
