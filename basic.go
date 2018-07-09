// 3.1 搭建一个web服务器

package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
	"html/template"
	"time"
	"crypto/md5"
	"io"
	"strconv"
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
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			// 验证token合法性
		} else {
			// 不存在token报错
		}
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
	}
}

// 默认情况下，Handler里面是不会自动解析form的，必须显式的调用r.ParseForm()后，
// 你才能对这个表单数据进行操作。我们修改一下代码，
// 在fmt.Println("username:", r.Form["username"])之前加一行r.ParseForm(),重新编译

func main(){
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil{
		log.Fatal("ListenAndServe: ", err)
	}
}