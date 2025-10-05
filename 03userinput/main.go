package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the userinput"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter that value :")

	//comma ok || err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Entered Value is : ", input)
}
