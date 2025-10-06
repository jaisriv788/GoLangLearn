package main

import "fmt"

func main() {
	fmt.Println("Welcome to if-else.")

	loginCount := -1

	var result string

	if loginCount >= 0 && loginCount <= 10 {
		result = "Regular User."
	} else if loginCount > 10 && loginCount <= 20 {
		result = "Silver User"
	} else if loginCount > 20 {
		result = "Gold User"
	} else {
		result = "Invalid Number"
	}

	fmt.Println(result)

	//Here the initialization and the checking is done in continuation

	if num := 10; num < 10 {
		fmt.Println("Result is less than 10.")
	} else {
		fmt.Println("Result is greater that 10.")
	}
}
