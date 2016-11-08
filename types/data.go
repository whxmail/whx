package types

import (
	"errors"
)

//Command line length
const (
	CMD_LENGTH = 10
)

type Data map[string]interface{}

//Flag is the type of string whitch contains some of symbols
func (data Data) GetFlag() (flag string, err error) {
	if v, ok := data["Flag"]; ok {
		if v2, ok := v.(string); ok {
			flag = v2
		} else {
			err = errors.New("Flag error!")
		}
	} else {
		err = errors.New("Can't find Flag!")
	}
	return flag, err
}

//CMD is the type of []string which contains a series of commands
func (data Data) GetCMD() (cmd []string, err error) {
	command := make([]string, CMD_LENGTH)
	length := 0
	if v, ok := data["CMD"]; ok {
		if v2, ok2 := v.([]interface{}); ok2 {
			for k, v3 := range v2 {
				var ok3 bool
				length++
				if length > CMD_LENGTH {
					err = errors.New("Command too long to control!")
					break
				}
				if command[k], ok3 = v3.(string); !ok3 {
					err = errors.New("Unknown command!")
				}
			}
		} else {
			err = errors.New("Command error!")
		}
	} else {
		err = errors.New("Can't find command!")
	}
	cmd = command[:length]
	return cmd, err
}
