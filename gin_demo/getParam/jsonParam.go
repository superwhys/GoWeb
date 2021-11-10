package getParam

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonParam(router *gin.Engine) {
	router.POST("/json", func(context *gin.Context) {
		b, _ := context.GetRawData()
		var m map[string]interface{}
		_ = json.Unmarshal(b, &m)

		context.JSON(http.StatusOK, m)
	})
}
