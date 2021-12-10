package main

import "fmt"

func main() {
	names := []string{"홍길동", "이순신", "감강찬"}
	for idx, name := range names {
		fmt.Println(idx, name)
	}

	var a = 1
	for a < 20 {
		println(a)
		if a == 5 {
			a += a
			continue
		}
		a += 2
		if a > 10 {
			break
		}
	}
	println(a)
	if a == 11 {
		goto END
	}
	println(a)
END:
	println("End")
}
