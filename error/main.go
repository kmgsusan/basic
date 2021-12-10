package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("C:\\temp\\error2.log")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(f.Name())
	switch err.(type) {
	default:
		println("ok")
	case error:
		log.Fatal(err.Error())
	}
}
