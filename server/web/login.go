package main

import (
	//"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
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
		usr := User{Username: r.Form.Get("username"), Password: r.Form.Get("password")}
		fmt.Println(usr)

		//ipclient
		dialer := "127.0.0.1"
		listener := "127.0.1.1"
		ic := ipClient{}
		ic.setAddr(dialer, listener)
		ic.startClient()
		//checkError(err)

		ic.login(usr)
		/*
			u, err := json.Marshal(usr)
			checkError(err)
			ic.conn.Write(u)

			//_, err = conn.Write([]byte(usr.password))
			//checkError(err)
		*/
		buf := make([]byte, 512)
		n, addr, err := ic.conn.ReadFromIP(buf)
		checkError(err)
		fmt.Println("Reply from server", addr.String(), string(buf[0:n]))

	}
}
