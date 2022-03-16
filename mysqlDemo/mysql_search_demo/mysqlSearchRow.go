package mysqlSearch

import (
	"database/sql"
	"fmt"
	"github.com/superwhys/GoWeb/mysqlDemo/mysqlInit"
)

func QueryRowDemo(sqlStr string, db *sql.DB) {
	var u mysqlinit.User
	err := db.QueryRow(sqlStr, 1).Scan(&u.Id, &u.Name, &u.Age)
	if err != nil {
		fmt.Printf("scan failed, err: %v\n", err)
	}

	fmt.Printf("id: %d, name: %s, age: %d\n", u.Id, u.Name, u.Age)

}
