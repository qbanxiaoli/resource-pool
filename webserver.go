package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println("web server")
	webserver := gin.Default()
	webserver.StaticFS("/public", http.Dir("/go_code/demo2"))
	webserver.GET("/getUsers", func(context *gin.Context) {
		fmt.Println("web server1")
		name :=  context.Query("name")
		password  := context.Query("password")
		fmt.Println("name:", name, "\tpassword:", password)
		context.JSON(200, gin.H{
			"user": "Jim",
		})
	})

	webserver.POST("/getUsers", func(context *gin.Context) {
		fmt.Println("web server1")
		// name :=  context.Query("name")
		// password  := context.Query("password")
		name := context.Request.FormValue("userName")
		password := context.Request.FormValue("pass")
		fmt.Println("name:", name, "\tpassword:", password)
		context.JSON(200, gin.H{
			"user": "Jim",
		})
	})
	_ = webserver.Run(":80")
}