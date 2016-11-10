// tcpServer
package types

import (
	"encoding/json"
	"fmt"
	"net"
)

var (
	TCP_DATA_FLAG chan int //It will recieved a int sum once TCPServer get data
)

func init() {
	TCP_DATA_FLAG = make(chan int)
}

//Router
type MUX map[string]func(Data, *TCPServer)

type TCPServer struct {
	conn net.Conn
	Mux  MUX
}

func NewTCPServer() (server TCPServer) {
	server = TCPServer{}
	server.Mux = make(map[string]func(Data, *TCPServer))
	return server
}

func (server *TCPServer) MuxRegister(cmd string, handleFunc func(Data, *TCPServer)) {
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
		<-TCP_DATA_FLAG
	}
}

func (server TCPServer) GetData() Data {
	buf := make([]byte, 512)
	n, err := server.conn.Read(buf)
	TCP_DATA_FLAG <- 1
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
