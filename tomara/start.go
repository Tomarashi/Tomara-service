package tomara

import (
	"fmt"
	"tomara-service/tomara/server"
)

func Start(host string, port int) {
	app := server.GetServer()
	app.Run(fmt.Sprintf("%s:%d", host, port))
}
