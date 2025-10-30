package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to Race Condition.")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var score = []int{0}

	//Here we have to tell the compiler that we have to wait for 3 goroutines
	//we can do it in two ways either by using Add before each go routine or we can use it once before all go routines
	wg.Add(3)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Routine One")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Routine Two")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Routine Three")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println("Score:", score)
}
