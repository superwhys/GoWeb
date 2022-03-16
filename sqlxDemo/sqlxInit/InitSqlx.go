package sqlxInit

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB(dsn string) (err error) {
	// 也可以使用MustConnect, 连接不成功就会panic
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("Connect DB failed, err: %v\n", err)
		return err
	}
	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(10)
	return nil
}
