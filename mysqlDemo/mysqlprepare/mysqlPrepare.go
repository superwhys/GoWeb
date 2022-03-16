package mysqlprepare

import (
	"database/sql"
	"fmt"
	"github.com/superwhys/GoWeb/mysqlDemo/mysqlInit"
)

func MysqlPrepareDemo(sqlStr string, db *sql.DB) {
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err: %v\n", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err: %v\n", err)
	}
	defer rows.Close()
	// 循环读取结果集中的数据

	for rows.Next() {
		var u mysqlinit.User
		err := rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			fmt.Printf("scan faild, err: %v\n", err)
		}
		fmt.Printf("id: %d, name: %s, age: %d\n", u.Id, u.Name, u.Age)
	}

}
