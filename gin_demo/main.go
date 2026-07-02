package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//2解析模板
	t,err :=template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("parse template failed, err:", err)
		return
	}
	//3渲染模板
	err = t.Execute(w, "小王子,sdsada")
	if err != nil {
		fmt.Println("execute template failed, err:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
    err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
		return
	}
}