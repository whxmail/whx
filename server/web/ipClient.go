package main

import (
	"encoding/json"
	"net"
)

//ipclient
type ipClient struct {
	dialer   string
	listener string
	conn     *net.IPConn
}

//Set address
func (ic *ipClient) set(dialer string, listener string) {
	ic.dialer = dialer
	ic.listener = listener
}

//Start the cilent
func (ic *ipClient) startClient() {
	lAddr, err := net.ResolveIPAddr("ip", ic.dialer)
	checkError(err)
	rAddr, err := net.ResolveIPAddr("ip", ic.listener)
	checkError(err)
	ic.conn, err = net.DialIP("ip:4", lAddr, rAddr)
	checkError(err)
	//return ic.conn, err
}

//login
func (ic ipClient) login(usr User) {
	usr.CMD = []string{"LOGIN", usr.Username, usr.Password}
	u, err := json.Marshal(usr)
	checkError(err)
	ic.conn.Write(u)
}
