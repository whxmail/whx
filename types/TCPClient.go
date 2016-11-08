// TCPClient
package types

import (
	"encoding/json"
	"fmt"
	"net"
)

type TCPClient struct {
	conn *net.TCPConn
}

func (client *TCPClient) StartClient(laddr string, raddr string) {
	lAddr, err := net.ResolveTCPAddr("tcp", laddr)
	checkError(err)
	rAddr, err := net.ResolveTCPAddr("tcp", raddr)
	checkError(err)
	client.conn, err = net.DialTCP("tcp", lAddr, rAddr)
	checkError(err)
}

func (client TCPClient) Login(usr User) {
	usr.CMD = CMD{"LOGIN", usr.Username, usr.Password}
	data := Data{"CMD": usr.CMD}
	client.SendData(data)
}

func (client TCPClient) GetData() Data {
	buf := make([]byte, 512)
	n, err := client.conn.Read(buf)
	checkError(err)
	var o interface{}
	json.Unmarshal(buf[:n], &o)
	data, ok := o.(map[string]interface{})
	if !ok {
		fmt.Println("Unknown data!")
	}
	return data
}

//data type: struct or map or json
func (client TCPClient) SendData(o interface{}) {
	data, _ := json.Marshal(o)
	client.conn.Write(data)
}
func (client TCPClient) Close() {
	client.Close()
}
