// mapTest.go
package main

import (
	"errors"
	"reflect"
	//	"encoding/json"
	"fmt"
)

type strTest []string

type structTest struct {
	strTest
}

type Data map[string]interface{}

func (data Data) GetValueByKey(key string) (value reflect.Value, err error) {
	if v, ok := data[key]; ok {
		value = reflect.ValueOf(v)
	} else {
		err = errors.New("err")
	}
	return value, err
}

func pt(k chan int) {
	<-k
	fmt.Println("go running...")
}

func main() {
	var k = make(chan int)

	for {
		go pt(k)
		k <- 1
	}

	/*
		data := Data{"name": "liu"}
		res, _ := data.GetValueByKey("name")
		//	var t map[string]interface{}
		//	t = data
		//	rest := reflect.TypeOf(res)
		datat := reflect.TypeOf(data)
		data_name := data["name"]
		data_name_t := reflect.TypeOf(data_name)
		fmt.Println(res)
		fmt.Println(data_name_t)
		fmt.Println(datat)
		/*
			var st structTest
			str := []string{"hello", "word"}
			st.strTest = str
			fmt.Println(str)
			fmt.Println(st)
	*/
	//map test
	/*	m := make(map[string]interface{})
			m = map[string]interface{}{"liu": 123, "zhu": []interface{}{"22", "33"}}
			b, _ := json.Marshal(m)
			fmt.Println(string(b))
			parseMap(m)
		}

		func parseMap(m map[string]interface{}) {
			for k, v := range m {
				switch vv := v.(type) {
				case string:
					fmt.Println(k, vv)
				case []string:
					for k2, v2 := range vv {
						fmt.Println(k2, v2)
					}
				case int:
					fmt.Println(k, vv)
				case []int:
					for k2, v2 := range vv {
						fmt.Println(k2, v2)
					}
				case map[string]interface{}:
					fmt.Println(k)
					parseMap(vv)
				default:
					fmt.Println("None support type!")

				}
			}
	*/
}
