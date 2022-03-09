package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"status": "OK",
		//})
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	r.GET("/a", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "a",
		})
	})

	r.GET("/b", func(c *gin.Context) {
		// 路由重定向
		// 把请求url修改
		c.Request.URL.Path = "/a"
		// 继续后续的处理
		r.HandleContext(c)
	})

	_ = r.Run(":9090")
}
