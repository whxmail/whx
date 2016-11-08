package types

import (
	"encoding/json"
	"fmt"
	"net"
)

//ipclient
type IPClient struct {
	lAddr string
	rAddr string
	conn  *net.IPConn
}

//Set address
func (ic *IPClient) SetAddr(laddr string, raddr string) {
	ic.lAddr = laddr
	ic.rAddr = raddr
}

//Start the cilent
func (ic *IPClient) StartClient() {
	laddr, err := net.ResolveIPAddr("ip", ic.lAddr)
	checkError(err)
	raddr, err := net.ResolveIPAddr("ip", ic.rAddr)
	checkError(err)
	ic.conn, err = net.DialIP("ip:4", laddr, raddr)
	checkError(err)
	//return ic.conn, err

}

//login
func (ic IPClient) Login(usr User) {
	usr.CMD = []string{"LOGIN", usr.Mail.String(), usr.Password}
	u, err := json.Marshal(usr)
	checkError(err)
	ic.conn.Write(u)
}

func (ic IPClient) GetData() Data {
	buf := make([]byte, 512)
	n, _, err := ic.conn.ReadFromIP(buf)
	checkError(err)
	//test(string(buf[:n]))
	var in interface{}
	json.Unmarshal(buf[:n], &in)
	//test(in)
	data, ok := in.(map[string]interface{})
	if !ok {
		fmt.Println("Unknown data!")
	}
	return data
}

//Data should be the type of struct or map
func (ic IPClient) SendData(data interface{}) {
	buf, _ := json.Marshal(data)
	//test(string(buf))
	ic.conn.Write(buf)
}
