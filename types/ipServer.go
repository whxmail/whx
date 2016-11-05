package types

import (
	"encoding/json"
	"fmt"
	//	"log"
	//	"fmt"
	//	"fmt"
	"net"
)

type IPServer struct {
	lAddr string
	conn  *net.IPConn
}

//设置ip监听地址
func (is *IPServer) SetAddr(laddr string) {
	is.lAddr = laddr
}

//启动ipServer
func (is *IPServer) StartServer() {
	ipAddr, err := net.ResolveIPAddr("ip", is.lAddr)
	checkError(err)
	is.conn, err = net.ListenIP("ip:4", ipAddr)
	checkError(err)
}

//Get data from IPClient
func (is IPServer) GetData() (Data, *net.IPAddr) {
	buf := make([]byte, 512)
	n, addr, err := is.conn.ReadFromIP(buf)
	checkError(err)
	var in interface{}
	json.Unmarshal(buf[:n], &in)
	//test(in)
	data, ok := in.(map[string]interface{})
	if !ok {
		fmt.Println("Unknown data!")
	}
	return data, addr
}

func (is IPServer) SendData(data interface{}, addr *net.IPAddr) {
	buf, err := json.Marshal(data)
	checkError(err)
	_, err = is.conn.WriteToIP(buf, addr)
	checkError(err)
}

func (is IPServer) GetConn() *net.IPConn {
	return is.conn
}
