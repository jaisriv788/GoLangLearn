package main

import (
	"fmt"
	"net/url"
)

const link = "https://jaisrivastava788.github.io:5173/cv?name=jai%20srivastava&age=25"

func main() {
	fmt.Println("Welcome To URL Handleing!")

	//Here is how to parse the url

	result, err := url.Parse(link)

	if err != nil {
		panic(err)
	}

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.RawQuery)
	// fmt.Println(result.Port())

	params := result.Query()
	fmt.Println(params)
	fmt.Println(params["name"])
	fmt.Println(params["age"])

	for i, val := range params {
		fmt.Printf("Key : %v; Value : %v \n", i, val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "http",
		Host:     "jaisrivastava788.github.io:5173",
		Path:     "/details",
		RawQuery: "email=jaisrivastava788@gmail.com&age=10",
	}

	fmt.Println(partsOfUrl)
}
