package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jaisrivastava788.github.io/cv/"

func main() {
	fmt.Println("Welcome to LCO Web Request.")

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("err")
		panic(err)
	}

	fmt.Printf("Type of Response : %T \n", response)

	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)
}
