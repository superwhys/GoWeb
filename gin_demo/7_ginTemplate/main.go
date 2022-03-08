package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 解析模板
	//r.LoadHTMLFiles("templates/index.html", "templates/user/index.html")
	r.LoadHTMLGlob("templates/**/*")

	r.GET("posts/index", func(c *gin.Context) {
		// 渲染模板
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "superyong post",
		})
	})

	r.GET("user/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/index.html", gin.H{
			"title": "superyong user",
		})
	})

	err := r.Run(":9000")
	if err != nil {
		fmt.Println(err)
		return
	}
}
