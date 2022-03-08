package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	UserName string `json:"username" form:"username"`
	PassWord string `json:"password" form:"password"`
}

func main() {
	router := gin.Default()

	router.GET("/user", func(c *gin.Context) {
		userName := c.Query("username")
		passWord := c.Query("password")
		// if there's a lot of data query data, ou have to write a lot of structures like this
		u := UserInfo{
			UserName: userName,
			PassWord: passWord,
		}
		fmt.Printf("%#v\n", u)
		c.JSON(http.StatusOK, gin.H{
			"username": userName,
			"password": passWord,
		})
	})

	// we can use bind arguments to improve above problem
	router.GET("/bindQuery", func(c *gin.Context) {
		// 声明一个UserInfo类型变量
		u := UserInfo{}
		// need to add tag: `form:"xxx"` in struct
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	router.POST("/bindForm", func(c *gin.Context) {
		// 声明一个UserInfo类型变量
		u := UserInfo{}
		// need to add tag: `form:"xxx"` in struct
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	router.POST("/bindJson", func(c *gin.Context) {
		// 声明一个UserInfo类型变量
		u := UserInfo{}
		// need to add tag: `form:"xxx"` in struct
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	err := router.Run(":9090")
	if err != nil {
		fmt.Println(err)
		return
	}
}
