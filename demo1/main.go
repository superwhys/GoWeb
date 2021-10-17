package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")	
	})

	err := r.Run()
	if err != nil {
		return 
	}
}

