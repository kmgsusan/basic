package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["a"] = "apple"
	m["g"] = "google"
	m["n"] = "naver"

	fmt.Println(m)
	if _, exist := m["c"]; !exist {
		fmt.Println("map null")
	} else {
		fmt.Println(m["c"])
	}

	// for 로 출력해보자
	for key, val := range m {
		fmt.Println(key, " => ", val)
	}
}
