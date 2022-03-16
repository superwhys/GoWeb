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

// QueryByIDs 根据给定ID查询
func QueryByIDs(db *sqlx.DB, ids []int) {
	// 动态填充id
	query, args, err := sqlx.In("SELECT name, age FROM user WHERE id IN (?)", ids)
	if err != nil {
		fmt.Println("in quert error: ", err)
		return
	}
	fmt.Println(query) // 查看生成的querystring
	fmt.Println(args)  // 查看生成的args
	// sqlx.In 返回带 `?` bindvar的查询语句, 我们使用Rebind()重新绑定它

	//query = db.Rebind(query)
	//fmt.Println(query) // 查看生成的querystring

	var users []User
	err = db.Select(&users, query, args...)
	if err != nil {
		fmt.Println("in select error: ", err)
		return
	}
	fmt.Println("user slice is: ", users)
}
