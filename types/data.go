package types

import (
	"errors"
	"log"
)

//Command line length
const ()

type Data map[string]interface{}
type Request Data
type Response Data

func (data Data) GetReq() (req string, err error) {
	v := data.GetValueByKey("REQ")
	req, ok := v.(string)
	if !ok {
		err = errors.New("map[\"REQ\"]: error type!")
	}
	return req, err
}

func (data Data) Status() (ok bool) {
	if _, ok := data["STATUS"]; ok {
		ok = false
	} else {
		ok = true
	}
	return ok
}

func (data Data) GetValueByKey(key string) (value interface{}) {
	value, ok := data[key]
	if !ok {
		log.Fatal("Can't find:" + key)
	}
	return value
}
