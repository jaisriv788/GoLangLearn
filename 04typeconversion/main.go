package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter the first value")
	value, _ := reader.ReadString('\n')

	fmt.Println("Entered first value is", value)
	fmt.Printf("Entered first value is of type %T \n", value)

	fmt.Printf("Enter the second value")
	var valueTwo int
	fmt.Scanln(&valueTwo)

	fmt.Println("Entered second value is", valueTwo)
	fmt.Printf("Entered second value is of type %T \n", valueTwo)

	convertedValue, error := strconv.ParseFloat(strings.TrimSpace(value), 64)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("The converted value is ", convertedValue)
		fmt.Printf("Entered value is of type %T", convertedValue)
	}
}
