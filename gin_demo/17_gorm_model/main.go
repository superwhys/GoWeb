package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`                            // 设置字段大小为255
	MemberNumber *string `gorm:"column:membernumber;unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`                      // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`                          // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`                                   // 忽略本字段
}

type Animal struct {
	AnimalId string `gorm:"PRIMARY_KEY"`
	Name     string
}

// TableName 可以自定义表名
func (Animal) TableName() string {
	return "test_model_name"
}

func main() {
	// GORM还支持更改默认表名称规则, 只会修改默认表名规则
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "prefix_" + defaultTableName
	}

	db, err := gorm.Open("mysql", "root@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	db.SingularTable(true)

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})

	// 也可以通过Table()指定表名：
	//db.Table("table_create_name").CreateTable(&User{})
}
