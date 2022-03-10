package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserTable struct {
	ID int64
	// tag gorm:"default:xxx" 设置默认值
	// 但对于该类型的零值，即""， 在数据库中不会写入该空字符串而会写入默认值
	// 若在设置了默认值时，想写入对应类型的零值，可以使用类型的指针, 或者实现 Scanner/Valuer接口
	Name    string         `gorm:"default:'defauleName'"`
	Gender  *string        `gorm:"default:'defaule_gender'"`
	Address sql.NullString `gorm:"default:'shenzhen'"` // sql.NullString 实现了Scanner/Valuer接口
	Age     int64
}

func main() {
	db, err := gorm.Open("mysql", "root@(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// 模型迁移，将模型与数据库表对应起来
	db.AutoMigrate(&UserTable{})

	// 创建
	u := UserTable{
		Gender: new(string),
		Address: sql.NullString{
			String: "",
			Valid:  true,
		},
		Age: 29,
	}
	// 判断主键是否为空
	fmt.Println(db.NewRecord(&u)) // true
	db.Create(&u)
	db.Debug().Create(&u)
	fmt.Println(db.NewRecord(&u)) // false
}
