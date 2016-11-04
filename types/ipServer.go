package types

import (
	"encoding/json"
	"fmt"
	"log"
	//	"fmt"
	//	"fmt"
	"net"
)

//type Data map[string]interface{}

//func(data Data) login

type IPServer struct {
	listener string
	conn     *net.IPConn
}

//设置ip监听地址
func (is *IPServer) SetAddr(listener string) {
	is.listener = listener
}

//启动ipServer
func (is *IPServer) StartServer() {
	ipAddr, err := net.ResolveIPAddr("ip", is.listener)
	checkError(err)
	is.conn, err = net.ListenIP("ip:4", ipAddr)
	checkError(err)
}

/*
func (is ipServer) GetUser() (usr User) {
	buf := make([]byte, 512)
	n, _, err := is.conn.ReadFromIP(buf)
	checkError(err)
	json.Unmarshal(buf[:n], &usr)
	return usr
}
*/
func (is IPServer) GetData() Data {
	buf := make([]byte, 512)
	n, _, err := is.conn.ReadFromIP(buf)
	checkError(err)
	var v interface{}
	json.Unmarshal(buf[:n], &v)
	//test(v)
	data, ok := v.(map[string]interface{})
	if !ok {
		fmt.Println("Unknown data!")
	}
	return data
}

func checkError(err error) {
	if err != nil {
		log.Fatal("Fatal error:", err.Error())
	}
}
