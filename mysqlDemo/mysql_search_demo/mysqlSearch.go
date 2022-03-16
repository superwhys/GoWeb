package mysqlSearch

import (
	"database/sql"
	"fmt"
	"github.com/superwhys/GinTest/mysqlDemo/mysqlInit"
)

func QueryMultiRowDemo(sqlStr string, db *sql.DB) {
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("quert faild, err: %v\n", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var u mysqlinit.User
		err = rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			fmt.Printf("scan faild, err: %v\n", err)
			return
		}

		fmt.Printf("id: %d, name: %s, age: %d\n", u.Id, u.Name, u.Age)
	}
}
