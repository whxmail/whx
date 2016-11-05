// funcs.go
package types

import (
	"fmt"
	"log"
)

func checkError(err error) {
	if err != nil {
		log.Fatal("Fatal error:", err.Error())
	}
}

func test(i interface{}) {
	fmt.Println("node:", i)
}
