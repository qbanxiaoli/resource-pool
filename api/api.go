package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"resource-pool/db"
	"strconv"
)

func Init(webserver *gin.Engine) {

	// 根据id查询服务器使用情况
	webserver.GET("/getResourceById", func(context *gin.Context) {
		id, err := strconv.Atoi(context.Query("id"))
		if err != nil {
			fmt.Println(err.Error())
		}
		context.JSON(http.StatusOK, db.GetResourceById(id))
	})

	// 查询资源池使用情况
	webserver.GET("/getResourceTotal", func(context *gin.Context) {
		currentPage := context.Query("currentPage")
		pageSize := context.Query("pageSize")
		if currentPage == "" {
			currentPage = "1"
		}
		if pageSize == "" {
			pageSize = "10"
		}
		page, err := strconv.Atoi(currentPage)
		if err != nil {
			fmt.Println(err.Error())
		}
		size, err := strconv.Atoi(pageSize)
		if err != nil {
			fmt.Println(err.Error())
		}
		context.JSON(http.StatusOK, db.GetResourceTotal(page, size))
	})

}
