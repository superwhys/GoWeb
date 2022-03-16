package main

import (
	"fmt"
	"github.com/superwhys/GinTest/sqlxDemo/sqlxDelete"
	"github.com/superwhys/GinTest/sqlxDemo/sqlxInit"
	"github.com/superwhys/GinTest/sqlxDemo/sqlxInsert"
	"github.com/superwhys/GinTest/sqlxDemo/sqlxNameExec"
	"github.com/superwhys/GinTest/sqlxDemo/sqlxNameQuery"
	"github.com/superwhys/GinTest/sqlxDemo/sqlxSearch"
	"github.com/superwhys/GinTest/sqlxDemo/sqlxUpdate"
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
	sqlxUpdate.UpdateRowDemo(sqlxInit.DB, 2)
	sqlxSearch.QueryRowDemo(sqlxInit.DB, 2)
	sqlxDelete.DeleteRowDemo(sqlxInit.DB, 2)
	sqlxNameExec.InsertUserDemo(sqlxInit.DB)
	sqlxNameQuery.NamedQuery(sqlxInit.DB)
}
