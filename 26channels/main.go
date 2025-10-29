package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to the channels")

	myCh := make(chan int)
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println("Outer", <-myCh)

	wg.Add(2)

	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-ch

		//if its closed isChannelOpen == flase if not then it will be true
		fmt.Println(val, isChannelOpen)
		fmt.Println(<-ch)
		wg.Done()
	}(myCh, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 5
		ch <- 51
		close(ch)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
