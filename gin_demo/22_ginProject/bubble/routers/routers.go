package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/superwhys/server/v1"
	"net/http"
)

var (
	Router *gin.Engine
)

func IndexRouter() {
	Router.Static("/static", "static")
	Router.LoadHTMLGlob("template/*")
	Router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}

func InitRouter() {
	Router = gin.Default()

	IndexRouter()

	// v1
	v1Group := Router.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("todo", v1.InsertHandler)
		// 查看
		// 查看所有
		v1Group.GET("todo", v1.SearchAllHandler)
		// 查看指定
		v1Group.GET("todo/:id", v1.SearchSpecifyHandler)
		// 修改
		v1Group.PUT("todo/:id", v1.UpdateHandler)
		// 删除
		v1Group.DELETE("todo/:id", v1.DeleteHandler)
	}
}
