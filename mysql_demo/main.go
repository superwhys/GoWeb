package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // init()
	"log"
)

var Db *sql.DB

func InitMysql(dsn string) (err error) {

	Db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	defer Db.Close()

	// 尝试与数据库建立连接 （校验dsn是否正确）
	err = Db.Ping()
	if err != nil {
		fmt.Println("connect to db faild, err: ", err)
		return
	}

	fmt.Println("connect to db success")
	return nil
}

func main () {
	dsn := "root:yang4869@tcp(localhost:3306)/test"
	if err := InitMysql(dsn); err != nil {
		log.Fatalln(err)
	}
}
