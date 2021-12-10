package main

// http://golang.site/go/article/13-Go-%EC%BB%AC%EB%A0%89%EC%85%98---Slice

import "fmt"

func main() {
	var a []int
	a = []int{1, 2, 3, 4}
	a[1] = 10
	fmt.Println(a)

	// Go에서 Slice를 생성하는 또 다른 방법으로 Go의 내장함수 make() 함수를 이용
	// 슬라이스를 make로 만들면 개발자가 임의로 length와 capacity를 임의로 조정해서 사용할 수 있다.
	s := make([]string, 3)
	fmt.Println(len(s), cap(s))
	fmt.Println(s)

	// nil slice
	var ss []int
	if ss == nil {
		println("Nil Slice")
	}
	fmt.Println(len(ss), cap(ss))

	// 3. 슬라이스 추가,병합(append)과 복사(copy)
	s2 := []int{0, 1}
	s2 = append(s2, 2)
	s2 = append(s2, 3, 4, 5, 6)
	fmt.Println(s2)
}
