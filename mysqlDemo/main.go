package main

import (
	"fmt"
	mysqlinit "github.com/superwhys/GoWeb/mysqlDemo/mysqlInit"
	mysqlSearch "github.com/superwhys/GoWeb/mysqlDemo/mysql_search_demo"
	"github.com/superwhys/GoWeb/mysqlDemo/mysqlprepare"
	"log"
)

func main() {

	//var db *sqlxIn.DB

	dsn := "root@(localhost:3306)/sql_test"

	if err := mysqlinit.InitMysql(dsn); err != nil {
		log.Fatalln(err)
	}

	db := mysqlinit.Db
	defer db.Close()

	fmt.Println("queryRow")
	sqlStr := "select id, name, age from user where id=?"
	mysqlSearch.QueryRowDemo(sqlStr, db)

	fmt.Println("queryMultiRow")
	fmt.Println("search more rows")
	sqlStr1 := "select id, name, age from user where id>?"
	mysqlSearch.QueryMultiRowDemo(sqlStr1, db)

	fmt.Println("MysqlPrepare")
	sqlStr2 := "select id, name, age from user where id>?"
	mysqlprepare.MysqlPrepareDemo(sqlStr2, db)
}
