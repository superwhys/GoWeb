package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Parse template
	// 在解析模板时，被嵌套的模板一定要在后面解析
	t, err := template.ParseFiles("./ol.html", "./ui.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Execute template
	name := "superyong"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
