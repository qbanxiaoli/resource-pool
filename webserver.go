package main

import (
	"github.com/gin-gonic/gin"
	"resource-pool/api"
	"resource-pool/config"
)

func main() {
	webserver := gin.Default()
	webserver.Use(config.Cors())
	api.Init(webserver)
	webserver.Run(":80")
}
