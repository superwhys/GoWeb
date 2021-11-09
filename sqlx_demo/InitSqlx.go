package sqlx_demo

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitDB(dsn string) (*sqlx.DB, error) {
	// 也可以使用MustConnect, 连接不成功就会panic
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("Connect DB failed, err: %v\n", err)
		return nil, err
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return db, nil
}
