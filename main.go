package main

import (
	"fmt"
	_ "gocloud/common/database"
	"gocloud/common/logger"
	"gocloud/routes"
	"net/http"
	"gocloud/common/config"
)

func main() {
	router := routes.Router()

	logger.Info("http server started, listened on port ", config.Config.Port)
	fmt.Printf("http server started, listened on port %s\n", config.Config.Port)
	logger.Fatal(http.ListenAndServe(":" + config.Config.Port, router))
}
