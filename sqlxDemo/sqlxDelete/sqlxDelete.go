package sqlxDelete

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// DeleteRowDemo 删除数据
func DeleteRowDemo(db *sqlx.DB, id int) {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}
