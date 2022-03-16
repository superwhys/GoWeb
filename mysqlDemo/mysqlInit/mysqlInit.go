package mysqlinit

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // init()
)

type User struct {
	Id   int    `db:"id"`
	Age  int    `db:"age"`
	Name string `db:"name"`
}

var Db *sql.DB

func InitMysql(dsn string) (err error) {

	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	//defer Db.Close()

	// 尝试与数据库建立连接 （校验dsn是否正确）
	err = Db.Ping()
	if err != nil {
		fmt.Println("connect to db faild, err: ", err)
		return
	}

	fmt.Println("connect to db success")
	Db.SetMaxOpenConns(100) // 最大连接数
	Db.SetMaxIdleConns(10)  // 最大空闲连接数
	return nil
}
