// interip
package types

import (
	"encoding/json"
	"fmt"
	"net"
)

type InterIP struct {
	lAddr    string
	rAddr    string
	listener *net.IPConn
	dialer   *net.IPConn
}

func (iip *InterIP) SetIP(laddr string, raddr string) {
	iip.lAddr = laddr
	iip.rAddr = raddr
}

func (iip *InterIP) StartInter() {
	laddr, err := net.ResolveIPAddr("ip", iip.lAddr)
	checkError(err)
	raddr, err := net.ResolveIPAddr("ip", iip.rAddr)
	checkError(err)
	iip.listener, err = net.ListenIP("ip:4", laddr)
	checkError(err)
	iip.dialer, err = net.DialIP("ip:4", laddr, raddr)
	checkError(err)
}

func (iip InterIP) GetData() Data {
	buf := make([]byte, 512)
	n, _, err := iip.dialer.ReadFromIP(buf)
	checkError(err)
	var in interface{}
	//test(string(buf[:n]))
	json.Unmarshal(buf[:n], &in)
	//test(in)
	data, ok := in.(map[string]interface{})
	if !ok {
		fmt.Println("Unknown data!")
	}
	return data

}

//data should be type map or struct
func (iip InterIP) SendData(data interface{}) {
	buf, err := json.Marshal(data)
	checkError(err)
	iip.dialer.Write(buf)

}
