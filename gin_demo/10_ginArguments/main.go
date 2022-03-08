package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// query Param /web?query=superyong
	router.GET("/web", func(context *gin.Context) {
		// 获取查询参数 quert string
		//name := context.Query("query")
		name := context.DefaultQuery("query", "yong")
		context.JSON(http.StatusOK, name)
	})

	// path Param /param/search/superyong/guangdong
	router.GET("/param/search/:username/:address", func(context *gin.Context) {
		// 获取路径参数 path string
		username := context.Param("username")
		address := context.Param("address")

		context.JSON(http.StatusOK, gin.H{
			"msg":      "OK",
			"username": username,
			"address":  address,
		})
	})

	// form Param
	router.POST("/param/search", func(context *gin.Context) {
		//username := context.PostForm("username")
		//address := context.PostForm("address")
		username := context.DefaultPostForm("username", "testname")
		address, ok := context.GetPostForm("address")
		if !ok {
			address = "tiantang"
		}

		context.JSON(http.StatusOK,
			gin.H{
				"msg":      "OK",
				"username": username,
				"address":  address,
			})
	})

	err := router.Run(":9090")
	if err != nil {
		fmt.Println(err)
	}
}
