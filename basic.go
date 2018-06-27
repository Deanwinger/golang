// 3.1 搭建一个web服务器

package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
	"html/template"
)

func sayHelloName(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello Alan")
}

func login(w http.ResponseWriter, r *http.Request){
	fmt.Println("method: ", r.method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
	}
}

func mian(){
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":9000", nil)
	if err != nil{
		log.Fatal("ListenAndServe: ", err)
	}
}