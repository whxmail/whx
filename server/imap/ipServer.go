package main

import (
	"encoding/json"
	"fmt"
	//	"fmt"
	//	"fmt"
	"net"
)

//type Data map[string]interface{}

//func(data Data) login

type ipServer struct {
	listener string
	conn     *net.IPConn
}

//设置ip监听地址
func (is *ipServer) set(listener string) {
	is.listener = listener
}

//启动ipServer
func (is *ipServer) startServer() {
	ipAddr, err := net.ResolveIPAddr("ip", is.listener)
	checkError(err)
	is.conn, err = net.ListenIP("ip:4", ipAddr)
	checkError(err)
}

/*
func (is ipServer) getUser() (usr User) {
	buf := make([]byte, 512)
	n, _, err := is.conn.ReadFromIP(buf)
	checkError(err)
	json.Unmarshal(buf[:n], &usr)
	return usr
}
*/
func (is ipServer) getData() map[string]interface{} {
	buf := make([]byte, 512)
	n, _, err := is.conn.ReadFromIP(buf)
	checkError(err)
	var v interface{}
	json.Unmarshal(buf[:n], &v)
	//test(v)
	m, ok := v.(map[string]interface{})
	if !ok {
		fmt.Println("Unknown data!")
	}
	return m
}
