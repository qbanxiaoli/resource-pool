package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resource-pool/db"
)

func Init(webserver *gin.Engine) {

	// 查询资源池使用情况
	webserver.GET("/getResourceTotal", func(context *gin.Context) {
		context.JSON(http.StatusOK, db.GetResourceTotal())
	})

}
