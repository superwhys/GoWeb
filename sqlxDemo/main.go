package main

import (
	"fmt"
	"github.com/superwhys/GoWeb/sqlxDemo/sqlxDelete"
	"github.com/superwhys/GoWeb/sqlxDemo/sqlxIn"
	"github.com/superwhys/GoWeb/sqlxDemo/sqlxInit"
	"github.com/superwhys/GoWeb/sqlxDemo/sqlxInsert"
	"github.com/superwhys/GoWeb/sqlxDemo/sqlxNameExec"
	"github.com/superwhys/GoWeb/sqlxDemo/sqlxNameQuery"
	"github.com/superwhys/GoWeb/sqlxDemo/sqlxSearch"
	"github.com/superwhys/GoWeb/sqlxDemo/sqlxUpdate"
)

func main() {
	dsn := "root@(localhost:3306)/sql_test"
	// ================== sqlx ==========================
	if err := sqlxInit.InitDB(dsn); err != nil {
		return
	}
	fmt.Println("init success")

	sqlxInsert.InsertRowDemo(sqlxInit.DB)
	sqlxSearch.QueryRowDemo(sqlxInit.DB, 1)
	sqlxSearch.QueryMultiRowDemo(sqlxInit.DB, 1)
	sqlxUpdate.UpdateRowDemo(sqlxInit.DB, 3)
	sqlxSearch.QueryRowDemo(sqlxInit.DB, 3)
	sqlxDelete.DeleteRowDemo(sqlxInit.DB, 5)
	sqlxNameExec.InsertUserDemo(sqlxInit.DB)
	sqlxNameQuery.NamedQuery(sqlxInit.DB)

	u1 := sqlxIn.User{Name: "wangwu", Age: 20}
	u2 := sqlxIn.User{Name: "wangwu2", Age: 23}
	users := []interface{}{
		u1,
		u2,
	}
	sqlxIn.BatchInsertUsers2(sqlxInit.DB, users)

	sqlxIn.QueryByIDs(sqlxInit.DB, []int{1, 3, 4})
}
