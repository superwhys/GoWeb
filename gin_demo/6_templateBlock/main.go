package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "this is index page"
	// 渲染模板
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, err := template.ParseFiles("./home.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "this is home page"
	// 渲染模板
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func index2(w http.ResponseWriter, r *http.Request) {
	// 定义模板（使用模板继承）
	// 解析模板
	t, err := template.ParseFiles("./template/base.tmpl", "./template/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "this is index2 page"
	// 渲染模板
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func home2(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, err := template.ParseFiles("./template/base.tmpl", "./template/home.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "this is home2 page"
	// 渲染模板
	//err = t.ExecuteTemplate(w, "home.html", msg)
	err = t.Execute(w, msg)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
