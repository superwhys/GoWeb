package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type user struct {
	Name   string
	gender string
	Age    int
}

func SayHello(w http.ResponseWriter, r *http.Request) {
	// 解析模版
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 渲染模版
	name := "Superyong"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func SayHello2(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 渲染模板
	u1 := user{
		Name:   "Superyong",
		gender: "man",
		Age:    18,
	}

	m1 := map[string]string{
		"name":   "superyong2",
		"gender": "male",
		"age":    "20",
	}

	hobbyList := []string{"java", "python", "Golang"}

	//err = t.Execute(w, u1)
	// if want to use u1 and m1 together
	err = t.Execute(w, map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"hobby": hobbyList,
	})

	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/", SayHello2)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
