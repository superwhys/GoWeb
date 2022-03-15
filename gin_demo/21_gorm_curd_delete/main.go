package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/superwhys/superLog"
)

type UserDelete struct {
	gorm.Model
	Name   string
	Age    int64
	Gender string
}

func main() {
	db, err := gorm.Open("mysql", "root@(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		superLog.Error(err)
		return
	}

	defer db.Close()

	db.AutoMigrate(&UserDelete{})

	//// 创建数据
	//u1 := &UserDelete{
	//	Name:   "superyong",
	//	Age:    18,
	//	Gender: "male",
	//}
	//db.Create(&u1)
	//u2 := &UserDelete{Name: "yanghaowen", Age: 20, Gender: "female"}
	//db.Create(&u2)

	// 删除
	// 软删除 设置deleted_at
	var user UserDelete
	user.ID = 1
	db.Debug().Delete(&user)

	// 删除匹配字段
	var user2 UserDelete

	db.Debug().Where("name=?", "yanghaowen").Delete(&user2)
	//db.Debug().Delete(UserDelete{}, "age=?", 21)

	// Unscoped 方法可以查询被软删除的记录
	var users []UserDelete
	db.Debug().Unscoped().Where("name = ?", "yanghaowen").Find(&users)
	superLog.Info(users)

	// 物理删除
	db.Debug().Unscoped().Where("name=?", "yanghaowen2").Delete(UserDelete{})
}
