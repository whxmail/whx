// handleClient.go
package main

import (
	"fmt"

	"github.com/whxmail/whx/types"
)

func handleClient(is types.IPServer) {
	data, addr := is.GetData() //map[string]interface{}
	cmd, err := data.GetCMD()
	if err != nil {
		resp := types.Data{"RESP": err}
		is.SendData(resp, addr)
	} else {
		switch cmd[0] {
		case "LOGIN":
		case "LOGOUT":
		}
	}
	fmt.Println("Hello World!")
}
