package sqlxNameExec

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func InsertUserDemo(db *sqlx.DB) {
	sqlStr := "INSERT INTO user (name,age) VALUES (:name,:age)"
	_, err := db.NamedExec(sqlStr,
		map[string]interface{}{
			"name": "super haowen",
			"age":  28,
		})
	if err != nil {
		fmt.Println("Name exec error: ", err)
		return
	}
}
