// mapTest.go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	m := make(map[string]interface{})
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

}
