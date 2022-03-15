package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/superwhys/database"
	"github.com/superwhys/model"
	"github.com/superwhys/superLog"
	"net/http"
)

func InsertHandler(c *gin.Context) {
	// 1. 从请求中把数据拿出来
	todoReq := model.Todo{}
	err := c.BindJSON(&todoReq)
	if err != nil {
		superLog.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 2. 存入数据库
	err = model.CreateTodo(&todoReq)
	if err != nil {
		superLog.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		// 3.返回响应
		c.JSON(http.StatusOK, todoReq)
	}
}

func SearchAllHandler(c *gin.Context) {
	var todoList []model.Todo
	if err := model.GetAllTodoList(&todoList); err != nil {
		superLog.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func SearchSpecifyHandler(c *gin.Context) {

}

func UpdateHandler(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "error id"})
		return
	}
	todo := model.Todo{}
	if err := database.DB.Where("id=?", id).First(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	// 获取传过来的状态
	c.BindJSON(&todo)
	if err := model.UpdateTodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteHandler(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "error id"})
		return
	}

	if err := model.DeleteTodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
