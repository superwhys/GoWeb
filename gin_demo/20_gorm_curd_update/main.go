package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/superwhys/superLog"
)

type UserTabel2 struct {
	gorm.Model
	Name   string
	Age    int64
	Active bool
}

func main() {
	db, err := gorm.Open("mysql", "root@(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		superLog.Error(err)
		return
	}
	defer db.Close()

	// 常见数据
	//db.AutoMigrate(&UserTabel2{})
	//
	//u1 := &UserTabel2{
	//	Name:   "superyong",
	//	Age:    18,
	//	Active: true,
	//}
	//db.Create(&u1)
	//
	//u2 := &UserTabel2{
	//	Name:   "yanghaowen",
	//	Age:    20,
	//	Active: false,
	//}
	//db.Create(&u2)

	// 查询
	var user UserTabel2
	db.First(&user)
	superLog.Info(user)

	// 更新
	user.Name = "superyong2"
	user.Age = 99

	db.Save(&user) // 默认会修改所有字段

	// 更新指定字段
	db.Debug().Model(&user).Update("name", "king")

	// 更新多个字段
	var user2 UserTabel2
	db.First(&user2)
	superLog.Info(user2)

	m := map[string]interface{}{
		"name":   "haowen yang",
		"age":    50,
		"active": true,
	}

	db.Debug().Model(&user2).Updates(m)
	m["age"] = 60
	db.Debug().Model(&user2).Select("age").Updates(m) // 只更新age字段
	m["name"] = "yang haowen"
	db.Debug().Model(&user2).Omit("active").Updates(m) // 忽略active字段

	superLog.Info("UpdateColumn")
	var user3 UserTabel2
	db.First(&user3)
	superLog.Info(user3)
	// 无hook更新，不会更新update_at字段
	// 更新单个属性，类似于 `Update`
	db.Debug().Model(&user3).UpdateColumn("name", "hello")

	// 更新多个属性，类似于 `Updates`
	db.Debug().Model(&user3).UpdateColumns(UserTabel2{Name: "hello", Age: 18})

	// 让User表中用户年龄在原来基础上+2
	var user4 UserTabel2
	db.First(&user4)
	superLog.Info(user4)

	db.Debug().Model(&user4).Update("age", gorm.Expr("age+?", 2))
	//db.Model(&user4).Update(map[string]interface{}{"age": gorm.Expr("age+?", 2)})

}
