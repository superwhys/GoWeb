package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})

	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "video index",
			})
		})
		videoGroup.POST("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "video index POST",
			})
		})
	}

	_ = r.Run(":9090")
}
