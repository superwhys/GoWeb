package main

import (
	"database/sql"
	"fmt"
	"gitee.com/superwhys/GinTest/mysqlInit"
	mysqlSearch "gitee.com/superwhys/GinTest/mysql_search_demo"
	"log"
)

func main() {

	var db *sql.DB

	dsn := "root:yang4869@tcp(localhost:3306)/sql_test"

	if err := mysqlinit.InitMysql(dsn); err != nil {
		log.Fatalln(err)
	}

	db = mysqlinit.Db
	defer db.Close()


	sqlStr := "select id, name, age from user where id=?"
	mysqlSearch.QueryRowDemo(sqlStr, db)
	
	fmt.Println("search more rows")
	sqlStr1 := "select id, name, age from user where id>?"
	mysqlSearch.QueryMultiRowDemo(sqlStr1, db)

}
