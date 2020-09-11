package main

import (
	"github.com/gin-gonic/gin"
	"resource-pool/api"
)

func main() {
	webserver := gin.Default()
	api.Init(webserver)
	webserver.Run(":80")
}
