package main

import "fmt"

const c int = 0
const s string = "hi"
const (
	Visa   = "Visa"
	Master = "MasterCard"
	Amex   = "American Express"
)

func main() {
	var f float32 = 11
	fmt.Println(f)
	var i, j, k int
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	rawLiteral := `
		아리랑\n
	아리랑\n
아라리요
	`
	interLiteral := "아리랑아리랑\n아라리요"
	fmt.Println(rawLiteral)
	fmt.Println(interLiteral)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2)
}
