package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/superwhys/database"
	"github.com/superwhys/model"
	"github.com/superwhys/routers"
	"github.com/superwhys/superLog"
)

func main() {
	// 连接数据库
	err := database.InitMysql()
	if err != nil {
		superLog.Error(err)
		return
	}
	defer database.DB.Close()
	// 创建数据表
	database.DB.AutoMigrate(&model.Todo{})

	routers.InitRouter()
	routers.Router.Run(":8080")
}
