package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/superwhys/superWeb/logger"
)

func SetUp() *gin.Engine {
	router := gin.New()
	router.Use(logger.GinLogger(), logger.GinRecovery(true))

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "OK",
		})
	})

	return router
}
