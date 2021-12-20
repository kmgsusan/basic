package main

import (
	"log"
	"os"
)

func main() {
	fname := "C:\\Users\\kmgsu\\go\\src\\github.com\\kmgsusan\\basic\\io\\test.txt"
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	byteSlice := make([]byte, 1212)
	bytesRead, err := file.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("os - Number of byte : %d \n", bytesRead)
	log.Printf("Data read : %s\n", byteSlice)
}
