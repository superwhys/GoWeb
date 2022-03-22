package mysql

import (
	"fmt"

	"github.com/superwhys/superLog"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var (
	Db *sqlx.DB
)

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	superLog.Info(dsn)
	Db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	Db.SetMaxOpenConns(viper.GetInt("mysql.max_connection"))
	Db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_connection"))
	return
}

func Close() {
	_ = Db.Close()
}
