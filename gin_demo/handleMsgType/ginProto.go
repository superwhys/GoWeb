package handleMsgtype

import (
	"gitee.com/superwhys/GinTest/gin_demo/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinProto(router *gin.Engine) {
	router.GET("/proto", func(context *gin.Context) {
		data := &proto.Test{
			Name:                 "superyong",
			Msg:                  "hello-world",
			Age:                  20,
		}
		context.ProtoBuf(http.StatusOK, data)
	})
}
