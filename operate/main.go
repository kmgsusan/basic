package main

import "fmt"

func main() {
	var k int = 10
	var p = &k
	println(*p)

	var i = 1
	var max = 3
	if val := i * 2; val < max {
		println(val)
	}
	i++
	println(i)

	var name string
	var category = 1
	switch category {
	case 1:
		name = "Paper Book"
	case 2:
		name = "eBook"
	case 3, 4:
		name = "Blog"
	default:
		name = "Other"
	}

	println(name)

	switch x := category << 3; x - 1 {
	default:
		println(x)
	}

	check(2)
	check2(2)
}

func check(val int) {
	switch val {
	case 1:
		fmt.Println("1 이하")
		fallthrough
	case 2:
		fmt.Println("2 이하")
		fallthrough
	case 3:
		fmt.Println("3 이하")
		fallthrough
	default:
		fmt.Println("default 도달")
	}
}

func check2(val int) {
	switch val {
	case 1:
		fmt.Println("1 이하")
	case 2:
		fmt.Println("2 이하")
	case 3:
		fmt.Println("3 이하")
	default:
		fmt.Println("default 도달")
	}
}
