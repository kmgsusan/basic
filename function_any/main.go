package main

type calculator func(int, int) int

func main() {
	sum := func(n ...int) int {
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}
	result := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	println(result)

	add := func(i int, j int) int {
		return i + j
	}

	result = calc(add, 20, 30)
	println(result)

	result = calc2(add, 10, 20)
	println(result)

}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

func calc2(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}
