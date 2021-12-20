package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	s := "hello world!"
	n, err := file.Write([]byte(s))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n, "바이트 저장완료")

	file2, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file2.Close()

	fi, err := file2.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	var data = make([]byte, fi.Size())
	n, err = file2.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n, "바이트 읽기완료")
	fmt.Println(string(data))

}
