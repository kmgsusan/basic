package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 123
	}()
	var i int
	i = <-ch
	println(i)

	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		done <- true
	}()

	fmt.Println("wait")
	<-done

	c := make(chan int, 1)
	c <- 1
	fmt.Println(<-c)
}
