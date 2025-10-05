package main

import "fmt"

func main() {
	fmt.Println("Welcome To Pointers")

	// var ptr *int
	// fmt.Println("The default value of a pointer is ", ptr)

	num := 10
	var numPtr = &num
	fmt.Println("The reference value of the pointer is - ", numPtr)
	fmt.Println("The value of the pointer is - ", *numPtr)

	*numPtr = *numPtr + 2
	fmt.Println("The value of the Integer after mutation is - ", num)
	fmt.Println("The value of the pointer after mutation is - ", *numPtr)

	num2 := num
	num2 = num2 + 2
	fmt.Println("The value of the Integer after mutation is - ", num)
	fmt.Println("The value of the Integer after mutation is - ", num2)

}
