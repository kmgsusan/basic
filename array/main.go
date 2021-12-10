package main

// http://golang.site/go/article/12-Go-%EC%BB%AC%EB%A0%89%EC%85%98---%EB%B0%B0%EC%97%B4

func main() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	println(a[2])

	var a1 = [3]int{1, 2, 3}
	var a3 = [...]int{4, 5, 6, 7}

	println(a1[0])
	println(a3[0])

	// 4. 다차원 배열의 초기화
	var aa = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	println(aa[0][1])
}
