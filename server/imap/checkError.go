package main

import (
	"log"
)

func checkError(err error) {
	if err != nil {
		log.Fatal("Fatal error:", err.Error())
	}
}
