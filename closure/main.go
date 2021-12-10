package main

func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	next := nextValue()
	println(next())
	println(next())
	println(next())

	next2 := nextValue()
	println(next2())
	println(next2())
}
