package sqlxUpdate

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// UpdateRowDemo 更新数据
func UpdateRowDemo(db *sqlx.DB, id int) {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db.Exec(sqlStr, 39, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}
