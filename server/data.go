package main

import (
	"fmt"
)

//	"encoding/json"
const (
	CMD_LENGTH = 10
)

type Data map[string]interface{}

func (data Data) parseData() {

}

func (data Data) getCMD() (cmd []string, length int) {
	command := make([]string, CMD_LENGTH)
	length = 0
	if v, ok := data["CMD"]; ok {
		if v2, ok2 := v.([]interface{}); ok2 {
			for k, v3 := range v2 {
				var ok3 bool
				length++
				if length > CMD_LENGTH {
					fmt.Println("Command must small than ", CMD_LENGTH)
					break
				}
				if command[k], ok3 = v3.(string); !ok3 {
					fmt.Println("Unknown command!")
				}
			}
		} else {
			fmt.Println("Command error!")
		}
	} else {
		fmt.Println("Can't find command!")
	}
	cmd = command[:length]
	return cmd, length
}
