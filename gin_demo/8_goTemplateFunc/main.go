package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()

	// gin template function
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	// 解析模板
	//r.LoadHTMLFiles("templates/index.html", "templates/user/index.html")
	r.LoadHTMLGlob("templates/**/*")

	r.GET("posts/index", func(c *gin.Context) {
		// 渲染模板
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "<a href='www.baidu.com'>百度</a>",
		})
	})

	err := r.Run(":9000")
	if err != nil {
		fmt.Println(err)
		return
	}
}
