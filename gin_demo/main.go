package main

import (
	"fmt"
	"gitee.com/superwhys/GinTest/gin_demo/getParam"
	"gitee.com/superwhys/GinTest/gin_demo/handleMsgType"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	handleMsgtype.GinJson(router)
	handleMsgtype.GinXML(router)
	handleMsgtype.GinYAML(router)
	handleMsgtype.GinProto(router)

	getParam.QueryString(router)
	getParam.FormParam(router)
	getParam.JsonParam(router)
	getParam.PathParam(router)
	getParam.BindParam(router)



	err := router.Run(":8000")
	if err != nil {
		fmt.Printf("router run err, %v\n", err)
		return
	}
}
