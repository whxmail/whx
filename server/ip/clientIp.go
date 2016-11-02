// ipClient.go
package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
/*
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s host", os.Args[0])
	}
	service := os.Args[1]
*/
	service := "127.0.0.1"
	lAddr, err := net.ResolveIPAddr("ip", service)
	checkError(err)
	name := "127.0.1.1"
	checkError(err)
	rAddr, err := net.ResolveIPAddr("ip", name)
	checkError(err)
	conn, err := net.DialIP("ip:4", lAddr, rAddr)
	checkError(err)
	_, err = conn.Write([]byte("Hello Server!"))
	checkError(err)
	var buf [512]byte
	n, addr, err := conn.ReadFromIP(buf[:])
	checkError(err)
	fmt.Println("Reply from server", addr.String(), string(buf[0:n]))
	//_, err = conn.WriteToIP([]byte("Hello Server!"),addr)
	//checkError(err)
	conn.Close()
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		log.Fatal("Fatal error:", err.Error())
	}
}
