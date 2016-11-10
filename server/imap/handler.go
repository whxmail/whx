// handler.go
package main

import (
	"fmt"

	"github.com/whxmail/whx/types"
)

func myHandleFunc(server *types.TCPServer) {
	data := server.GetData()
	req, err := data.GetReq()
	checkError(err)
	handleFunc, ok := server.Mux[req]
	if !ok {
		fmt.Println("Request not register!")
	}
	handleFunc(data, server)
	//server.Close()
}

func login(data types.Data, server *types.TCPServer) {
	resp := types.Data{"RESP": "Received",
		"Yourname:": data["Username"]}
	server.SendData(resp)
	test("Dealed!")
}
