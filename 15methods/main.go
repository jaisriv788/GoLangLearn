package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Welcome to methods")

	jai := User{"Jai Srivastava", "jaisrivastava788@gmail.com", true, 27}

	jai.GetStatus()
	jai.NewMail()
	fmt.Println("Original Email: ", jai.Email)
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
