package model

import "github.com/superwhys/database"

// Todo model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// TableName 可以自定义表名
func (Todo) TableName() string {
	return "bubble"
}

func CreateTodo(td *Todo) (err error) {
	err = database.DB.Create(&td).Error
	return
}

func GetAllTodoList(todoList *[]Todo) (err error) {
	err = database.DB.Find(todoList).Error
	return
}

func UpdateTodo(todo *Todo) (err error) {
	err = database.DB.Save(todo).Error
	return
}

func DeleteTodo(id string) (err error) {
	err = database.DB.Where("id=?", id).Delete(Todo{}).Error
	return
}
