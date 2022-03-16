package sqlxInsert

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// InsertRowDemo 插入数据
func InsertRowDemo(db *sqlx.DB) {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db.Exec(sqlStr, "zhangsan", 30)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}
