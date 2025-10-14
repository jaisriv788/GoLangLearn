package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to server")

	getRequest()
	postRequest()
	postFormRequest()
}

func getRequest() {
	const myurl = "http://localhost:3000/get"

	res, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Status code: ", res.StatusCode)
	fmt.Println("Length: ", res.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(res.Body)
	bytecount, _ := responseString.Write(content)

	fmt.Println("Bytecount is: ", bytecount)
	fmt.Println(responseString.String())

	fmt.Println(string(content))
}

func postRequest() {
	const myurl = "http://localhost:3000/get"

	requestBody := strings.NewReader(`
	{
	"name" : "Jai Srivastava",
	"age":"23"
	}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	bytecount, _ := responseString.Write(content)

	fmt.Println("Bytecount is: ", bytecount)
	fmt.Println(responseString.String())
}

func postFormRequest() {
	const myurl = "http://localhost:3000/get"

	//form data
	data := url.Values{}
	data.Add("name", "Jai Srivastava")
	data.Add("age", "25")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	content, _ := io.ReadAll(response.Body)

	defer response.Body.Close()

	fmt.Println(string(content))
}
