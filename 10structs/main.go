package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Welcome to struct")

	jai := User{"Jai Srivastava", "jaisrivastava788@gmail.com", true, 27}
	fmt.Println("This is the registered User ", jai)
	fmt.Printf("This is the registered User %+v\n", jai)
	fmt.Println("This is the registered User name ", jai.Name)
}
