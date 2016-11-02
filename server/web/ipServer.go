package main

import (
	"net"
)

type ipServer struct {
	listener string
}

//设置ip监听地址
func(s *ipServer) set(addr string) {
	s.listener = addr
}

//启动ipServer
func(s ipServer) startServer() (conn *net.IPConn,err error) {
	ipAddr, err := net.ResolveIPAddr("ip", s.listener)
	checkError(err)
	conn,err = net.ListenIP("ip:4", ipAddr)
	checkError(err)
	return conn,err
}

