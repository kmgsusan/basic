package main

import "fmt"

func main() {
	count, total, msg := sum(1, 2, 3, 4, 6, 10, 100)
	fmt.Println(count, total, msg)
}
func sum(nums ...int) (count int, total int, msg string) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	msg = "okok"
	return
}
