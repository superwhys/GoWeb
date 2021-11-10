package getParam

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// QueryString /user/search?username=小王子&address=沙河
func QueryString(router *gin.Engine) {
	router.GET("/param/search", func(context *gin.Context) {
		username := context.DefaultQuery("username", "why")
		address := context.Query("address")

		context.JSON(http.StatusOK,
			gin.H{
			"message": "OK",
			"username": username,
			"address": address,
			})
	})
}
