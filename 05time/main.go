package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of GoLang")

	currentDate := time.Now()
	fmt.Println("Current timestamp is = ", currentDate)
	fmt.Printf("Current timestamp is of type %T \n", currentDate)

	//format must be same
	fmt.Println("Time formatting = ", currentDate.Format("02-01-2006 15:04:05 Monday"))

	dateCreated := time.Date(2023, time.October, 28, 11, 11, 11, 11, time.UTC)

	fmt.Println("Create date is - ", dateCreated.Format("02-01-2006 15:04:05 Monday"))

	fmt.Println("\nPress Enter to exit...")
	fmt.Scanln()
}
