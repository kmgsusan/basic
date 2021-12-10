package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := person{}
	p.name = "mgkim"
	p.age = 40
	fmt.Println(p)

	var p1 person
	p1 = person{"Bob", 20}
	p2 := person{name: "selee", age: 37}
	fmt.Println(p1, p2)
}
