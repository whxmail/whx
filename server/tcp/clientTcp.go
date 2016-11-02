// client.go
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s host:port", os.Args[0])
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	addr := conn.RemoteAddr()
	n, err := conn.Write([]byte("Hello server!"))
	checkError(err)
	var buf [512]byte
	n, err = conn.Read(buf[:])
	checkError(err)
	fmt.Println("Reply from server", addr.String(), string(buf[0:n]))
	conn.Close()
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}
