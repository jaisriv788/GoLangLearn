package main

import "fmt"

func main() {
	fmt.Println("Welcome to array.")

	//Declaration
	var arr [4]string

	arr[0] = "Iron"
	arr[1] = "Copper"
	arr[2] = "Silver"
	arr[3] = "Gold"

	fmt.Println("Element are", arr)
	fmt.Println("Element lenght is", len(arr))

	//Second way to declare it
	var arr2 = [3]int{1, 2, 3}

	fmt.Println("Element are", arr2)
	fmt.Println("Element lenght is", len(arr2))
}
