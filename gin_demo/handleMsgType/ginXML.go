package handleMsgtype

import (
	"gitee.com/superwhys/GinTest/gin_demo/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinXML(router *gin.Engine) {
	router.GET("/someXML", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{"message": "hello world"})
	})

	router.GET("/moreXML", func(context *gin.Context) {
		msg := config.Msg{
			Name:    "superyong",
			Message: "hello world",
			Age:     18,
		}

		context.XML(http.StatusOK, msg)
	})
}
