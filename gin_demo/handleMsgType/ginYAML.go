package handleMsgtype

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinYAML(router *gin.Engine) {
	router.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "ok", "status": http.StatusOK})
	})
}
