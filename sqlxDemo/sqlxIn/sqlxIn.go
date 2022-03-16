package sqlxIn

import (
	"database/sql/driver"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Age  int    `db:"age"`
	Name string `db:"name"`
}

func (u User) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

// BatchInsertUsers2 使用sqlx.In帮我们拼接语句和参数, 注意传入的参数是[]interface{}
func BatchInsertUsers2(db *sqlx.DB, users []interface{}) {
	query, args, err := sqlx.In(
		"INSERT INTO user (name, age) VALUES (?), (?)",
		users..., // 如果arg实现了 driver.Valuer, sqlx.In 会通过调用 Value()来展开它
	)
	if err != nil {
		fmt.Println("In error: ", err.Error())
		return
	}
	fmt.Println(query) // 查看生成的querystring
	fmt.Println(args)  // 查看生成的args
	_, err = db.Exec(query, args...)
	if err != nil {
		fmt.Println("batch insert error", err.Error())
		return
	}
}
