package database

import "github.com/jinzhu/gorm"

var (
	DB *gorm.DB
)

func InitMysql() (err error) {
	dsn := "root@(localhost:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}
