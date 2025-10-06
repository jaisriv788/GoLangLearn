package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value On The Dice Is = ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Move 1 step")
	case 2:
		fmt.Println("Move 2 step")
	case 3:
		fmt.Println("Move 3 step")
	case 4:
		fmt.Println("Move 4 step")
		fallthrough //this will run the statements in the next cases also.
	case 5:
		fmt.Println("Move 5 step")
		fallthrough
	case 6:
		fmt.Println("Move 6 step")
	default:
		fmt.Println("Invalid Number")
	}
}
