package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
	}()
	go func() {
		for i := 0; i < 7; i++ {
			c <- i
		}
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}

	// http://pyrasis.com/book/GoForTheReallyImpatient/Unit34/04

	d := make(chan int)
	go producer(d)
	go consumer(d)

}

func producer(c chan<- int) { // 보내기 전용
	for i := 0; i < 5; i++ {
		c <- i
	}
	c <- 100
}

func consumer(c <-chan int) { // 받기전용
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println(<-c)
}
