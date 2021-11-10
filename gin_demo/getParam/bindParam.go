package getParam

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func BindParam(router *gin.Engine) {
	router.POST("/bindParam", func(context *gin.Context) {
		var login Login
		if err := context.ShouldBind(&login); err == nil {
			fmt.Printf("login info: %v\n", login)
			context.JSON(http.StatusOK, gin.H{
				"username": login.User,
				"password": login.Password,
			})
		} else {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	router.GET("/bindParam", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
}
