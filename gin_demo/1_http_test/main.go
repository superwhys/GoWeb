package main

import (
	"fmt"
	"net/http"
)

// http server

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world！")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9915", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
