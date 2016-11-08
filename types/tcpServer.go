// tcpServer
package types

import (
	"encoding/json"
	"fmt"
	"net"
)

//Router
type MUX map[string]func([]string, *TCPServer)

type TCPServer struct {
	conn net.Conn
	Mux  MUX
}

func NewTCPServer() (server TCPServer) {
	server = TCPServer{}
	server.Mux = make(map[string]func([]string, *TCPServer))
	return server
}

func (server *TCPServer) MuxRegister(cmd string, handleFunc func([]string, *TCPServer)) {
	server.Mux[cmd] = handleFunc
}

func (server *TCPServer) ListenAndServer(tcpaddr string, handleFunc func(*TCPServer)) {
	var err error
	tcpAddr, err := net.ResolveTCPAddr("tcp", tcpaddr)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	server.conn, _ = listener.Accept()
	for {

		go handleFunc(server)
	}
}

func (server TCPServer) GetData() Data {
	buf := make([]byte, 512)
	n, err := server.conn.Read(buf)
	checkError(err)
	var o interface{}
	json.Unmarshal(buf[:n], &o)
	data, ok := o.(map[string]interface{})
	if !ok {
		fmt.Println("Unknown data!")
	}
	return data
}

//data that you want to send must be the type of struct or map or json
func (server TCPServer) SendData(o interface{}) {
	data, _ := json.Marshal(o)
	server.conn.Write(data)
}

func (server TCPServer) Close() {
	server.conn.Close()
}

/*
func (ts TCPServer) SetAddr(laddr string) {
	ts.lAddr = laddr
}

func (ts TCPServer) StartListener() {
	laddr, err := net.ResolveTCPAddr("tcp", ts.lAddr)
	checkError(err)
	ts.listener, err = net.ListenTCP("tcp", laddr)
	checkError(err)
}
*/
