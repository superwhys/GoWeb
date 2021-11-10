package handleMsgtype

import (
	"gitee.com/superwhys/GinTest/gin_demo/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinJson(router *gin.Engine) {
	router.GET("/someJSON", func(context *gin.Context) {
		// 方式一：自己拼接JSON
		context.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
	})

	router.GET("/moreJSON", func(context *gin.Context) {
		// 方法二：使用结构体
		msg := config.Msg{}
		msg.Name = "小王子"
		msg.Message = "Hello world!"
		msg.Age = 18
		context.JSON(http.StatusOK, msg)
	})
}
