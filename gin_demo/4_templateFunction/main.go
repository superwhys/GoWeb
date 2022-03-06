package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Handle1(w http.ResponseWriter, r *http.Request) {
	// 定义一个函数kua, 需在解析模板之前添加函数
	// 要么只有一个返回值，要么有两个返回值，第二个必须时error类型
	kua := func(name string) (string, error) {
		return name + " so best", nil
	}
	// 解析模版
	// 使用New创建模版时，名字一定要与filename对得上
	// Funcs 给模板添加函数
	t, err := template.New("tempDemo.html").Funcs(template.FuncMap{
		"kua": kua,
	}).ParseFiles("./tempDemo.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	name := "superyong"
	// 渲染模板
	err = t.Execute(w, map[string]interface{}{
		"name": name,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/", Handle1)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
