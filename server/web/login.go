package main

import (
	//"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/whxmail/whx/types"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm()
		usr := types.User{Username: r.Form.Get("username"), Password: r.Form.Get("password")}
		fmt.Println(usr)

		client.Login(usr)

		data := client.GetData()
		fmt.Println(data)
		//	client.Close()
	}
}
