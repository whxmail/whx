package main

import (
	"net"
)

type ipClient struct {
	dialer string
	listener string
}

func(c *ipClient) set(d string,l string) {
	c.dialer = d
	c.listener = l
}

func(c ipClient) startClient() (conn*net.IPConn,err error){
	lAddr, err := net.ResolveIPAddr("ip", c.dialer)
	checkError(err)
	rAddr, err := net.ResolveIPAddr("ip", c.listener)
	checkError(err)
	conn, err = net.DialIP("ip:4", lAddr, rAddr)
	checkError(err)
	return conn,err
}
