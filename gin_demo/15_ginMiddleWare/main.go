package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func indexHandle(c *gin.Context) {
	fmt.Println("this is index")
	//name := c.MustGet("name")
	name, ok := c.Get("name")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "name not exists",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

func middleWareHandle(c *gin.Context) {
	fmt.Println("this is middle")
	c.Next()
	fmt.Println("this is middle out")
}

func middleWareHandle2(c *gin.Context) {
	fmt.Println("this is middle2")
	// 在gin.context中存储值，以便被其它handle使用
	c.Set("name", "superyong")
	//c.Next()
	fmt.Println("this is middle2 out")
}

func middleWareHandle3(c *gin.Context) {
	fmt.Println("this is middle3")
	c.Abort()
	fmt.Println("this is middle3 out")
}

func timeCost(c *gin.Context) {
	fmt.Println("this is timeCost")
	// 调用后续的函数
	start := time.Now()

	// 在goroutine中使用context时，需要使用其只读副本
	//go func(c *gin.Context) {
	//	return
	//}(c.Copy())
	c.Next()
	cost := time.Since(start)
	fmt.Printf("cost: %v\n", cost)
	// 阻止调用后续的函数
	//c.Abort()
}

func main() {
	r := gin.Default()

	// 使用全局中间件
	r.Use(timeCost, middleWareHandle2)

	// 针对某一条路由使用中间件
	r.GET("/index", middleWareHandle, indexHandle)
	r.GET("/shop", indexHandle)

	// 路由组注册中间件
	testGroup := r.Group("/test", timeCost)
	{
		testGroup.GET("/index", indexHandle)
	}

	_ = r.Run(":9091")
}
