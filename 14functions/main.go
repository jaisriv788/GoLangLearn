package main

import "fmt"

func main() {
	fmt.Println("Welcome to the function.")
	sum := addition(2, 5)

	fmt.Println("Sum of the number is ", sum)

	supersum := proAddition(5, 4, 7, 3, 5, 7, 6, 4, 1, 2, 3, 7, 5, 4, 6)
	fmt.Println("Sum of the n number is ", supersum)
}

func addition(num1, num2 int) int {
	return num1 + num2
}

func proAddition(values ...int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}
