package getParam

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FormParam(router *gin.Engine) {
	router.POST("/param/search", func(context *gin.Context) {
		username := context.PostForm("username")
		address := context.PostForm("address")

		context.JSON(http.StatusOK,
			gin.H{
			"msg": "OK",
			"username": username,
			"address": address,
			})
	})
}
