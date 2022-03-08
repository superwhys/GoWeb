package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	router := gin.Default()

	router.LoadHTMLFiles("index.html")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/upload", func(c *gin.Context) {
		// 处理multipart forms提交文件时默认的内存限制是32 MiB
		// 可以通过下面的方式修改
		// router.MaxMultipartMemory = 8 << 20  // 8 MiB

		// 1. 读取文件 FormFile
		fileReader, err := c.FormFile("f1")
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			// 2. 将读取到的文件保存到本地
			//filePath := fmt.Sprintf("./%s", fileReader.Filename)
			filePath := path.Join("./", fileReader.Filename)
			_ = c.SaveUploadedFile(fileReader, filePath)
			c.JSON(http.StatusOK, gin.H{
				"fileName": fileReader.Filename,
			})
		}
	})

	err := router.Run(":9090")
	if err != nil {
		fmt.Println(err)
		return
	}
}
