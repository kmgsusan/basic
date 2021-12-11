package main

import (
	"fmt"
	"sync"
	"time"
)

var globalValue int

// func action(i int) {
// 	globalValue += i
// 	time.Sleep(1 * time.Second)
// }
// func action(i int, wg *sync.WaitGroup) {
// 	globalValue += i
// 	time.Sleep(1 * time.Second)
// 	wg.Done()
// }
func action(i int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	mutex.Lock()
	globalValue += i
	mutex.Unlock()
	time.Sleep(1 * time.Second)
	wg.Done()
}

// https://lynlab.co.kr/blog/82

func main() {
	startTime := time.Now()
	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(100)
	for i := 0; i < 100; i++ {
		// go action(i)
		go action(i, &mutex, &wg)
	}
	wg.Wait()
	delta := time.Now().Sub(startTime)
	fmt.Printf("Result is %d, done is %.3fs.\n", globalValue, delta.Seconds())
}
