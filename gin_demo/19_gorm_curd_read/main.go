package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/superwhys/superLog"
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

// main gorm 查询操作
func main() {
	db, err := gorm.Open("mysql", "root@(localhost:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		superLog.Error(err)
		return
	}

	var user UserTable
	// 根据主键查询第一条记录
	db.First(&user)
	//// SELECT * FROM users ORDER BY id LIMIT 1;
	superLog.Info(user)

	// 随机获取一条记录
	db.Take(&user)
	//// SELECT * FROM users LIMIT 1;
	superLog.Info(user)

	// 根据主键查询最后一条记录
	db.Last(&user)
	//// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	superLog.Info(user)

	var users []UserTable
	// 查询所有的记录
	db.Find(&users)
	//// SELECT * FROM users;
	superLog.Info(users)

	// 查询指定的某条记录(仅当主键为整型时可用)
	db.First(&user, 10)
	//// SELECT * FROM users WHERE id = 10;
	superLog.Info(user)

}
