// handler.go
package main

import (
	"fmt"
	//	"fmt"

	"github.com/whxmail/whx/types"
)

func myHandleFunc(server *types.TCPServer) {
	data := server.GetData()
	//server.SendData(types.Data{"hello": "client1"})
	cmd, err := data.GetCMD()
	//test(cmd)
	checkError(err)
	handleFunc, ok := server.Mux[cmd[0]]
	if !ok {
		fmt.Println("Command not register!")
	}
	handleFunc(cmd, server)
	//server.Close()
}

func login(cmd []string, server *types.TCPServer) {
	data := types.Data{"RESP": "Hello client2"}
	server.SendData(data)
	test("Dealed!")
}
