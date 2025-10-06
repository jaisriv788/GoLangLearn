package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer.")

	defer fmt.Println("Hello World")

	defer fmt.Println("Hi How are you?")

	myDefer()
}

//when ever a line have a defer keyword then it will be executed at the end of the code
//but if more than one lines are marked defered then the execution happens in reverse order
//i.e. the last defer will execute first and first defer in the program will execute at last.

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
