package sqlxSearch

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/superwhys/GoWeb/mysqlDemo/mysqlInit"
)

// QueryRowDemo 查询单条数据示例
func QueryRowDemo(db *sqlx.DB, id int) {
	sqlStr := "select id, name, age from user where id=?"
	var u mysqlinit.User
	err := db.Get(&u, sqlStr, id)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
}

// QueryMultiRowDemo 查询多条数据示例
func QueryMultiRowDemo(db *sqlx.DB, id int) {
	sqlStr := "select id, name, age from user where id > ?"
	var users []mysqlinit.User
	err := db.Select(&users, sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}
