package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {

	db, err := gorm.Open("mysql", "root@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// 创建表，自动迁移(把数据结构和数据表一一对应)
	db.AutoMigrate(&UserInfo{})

	// 创建数据行
	//u1 := UserInfo{1, "superyong", "male", "coding"}
	//db.Create(&u1)

	// 查询
	var u2 UserInfo
	db.First(&u2) // 查询表中第一条数据
	fmt.Printf("%#v\n", u2)

	// 更新
	db.Model(&u2).Update("hobby", "swimming")
	// 删除
	db.Delete(&u2)
}
