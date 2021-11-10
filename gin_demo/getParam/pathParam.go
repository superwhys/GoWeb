package getParam

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PathParam(router *gin.Engine) {
	router.GET("/param/search/:username/:address", func(context *gin.Context) {
		username := context.Param("username")
		address := context.Param("address")

		context.JSON(http.StatusOK, gin.H{
			"msg": "OK",
			"username": username,
			"address": address,
		})
	})
}
