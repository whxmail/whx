package main

import (
	"html/template"
	"fmt"
	"net/http"
	"log"
)

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //获取请求的方法
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        log.Println(t.Execute(w, nil))
    } else {
        //请求的是登陆数据，那么执行登陆的逻辑判断
	r.ParseForm()
        //fmt.Println("username:", r.Form["username"])
        //fmt.Println("password:", r.Form["password"])
	usr := user{username:r.Form.Get("username"),password:r.Form.Get("password")}
	fmt.Println(usr)
    }
}

type user struct {
	username string
	password string
}
