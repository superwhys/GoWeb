package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/json", func(context *gin.Context) {
		// 方法1：使用map
		//data := map[string]interface{}{
		//	"name": "superyong",
		//	"age":  18,
		//}
		// 2. gin.H 实际相当于方法1
		data := gin.H{
			"name": "superyong",
			"age":  18,
		}
		context.JSON(http.StatusOK, data)
	})

	router.GET("/struct_json", func(context *gin.Context) {
		// 3. struct
		type msg struct {
			Name string `json:"name"`
			Age  int
		}
		data := msg{
			Name: "superyong",
			Age:  18,
		}
		context.JSON(http.StatusOK, data)
	})

	err := router.Run(":9090")
	if err != nil {
		fmt.Println(err)
		return
	}
}
