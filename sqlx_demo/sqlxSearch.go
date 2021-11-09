package sqlx_demo

import (
	"fmt"
	mysqlinit "gitee.com/superwhys/GinTest/mysqlInit"
	"github.com/jmoiron/sqlx"
)

// 查询单条数据示例
func queryRowDemo(db *sqlx.DB) {
	sqlStr := "select id, name, age from user where id=?"
	var u mysqlinit.User
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
}

// 查询多条数据示例
func queryMultiRowDemo(db sqlx.DB) {
	sqlStr := "select id, name, age from user where id > ?"
	var users []mysqlinit.User
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}