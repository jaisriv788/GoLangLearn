package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to File Handling.")

	content := "This needs to be in a file - Jai Srivastava"

	file, err := os.Create("./dbFile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Length is ", length)

	defer file.Close()
	readFile("./dbFile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("Text data is \n", string(databyte))
}
