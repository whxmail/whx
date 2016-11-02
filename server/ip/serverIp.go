// ipServer
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//name, err := os.Hostname()
	//checkError(err)
	service := "127.0.1.1"
	ipAddr, err := net.ResolveIPAddr("ip", service)
	checkError(err)
	conn, err := net.ListenIP("ip:4", ipAddr)
	checkError(err)

	fmt.Println("1.sucess....")

	for {
		fmt.Println("2.sucess....")
		handleClient(conn)
		fmt.Println("3.sucess....")
	}
}

func handleClient(conn *net.IPConn) {
	var buf [512]byte
	n, addr, err := conn.ReadFromIP(buf[:])
	if err != nil {
		return
	}
	fmt.Println("Receive from client", addr.String(), string(buf[:n]))
	conn.WriteToIP([]byte("Welcome Client!"), addr)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}
