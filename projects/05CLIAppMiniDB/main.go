package main

import (
    "fmt"
    "strings"
)

// Person struct to store user data
type Person struct {
    Name   string
    Age    int
    Gender string
}

func main() {
    var people []Person // slice to store all entries
    for {
        fmt.Println("\n--- Person Management CLI ---")
        fmt.Println("1. Add Person")
        fmt.Println("2. Show All People")
        fmt.Println("3. Exit")
        fmt.Print("Choose an option: ")

        var choice int
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            var p Person
            fmt.Print("Enter Name: ")
            fmt.Scanln(&p.Name)
            fmt.Print("Enter Age: ")
            fmt.Scanln(&p.Age)
            fmt.Print("Enter Gender: ")
            fmt.Scanln(&p.Gender)

            people = append(people, p)
            fmt.Println("Person added successfully!")

        case 2:
            if len(people) == 0 {
                fmt.Println("No people stored yet.")
            } else {
                fmt.Println("\n--- People Table ---")
                fmt.Printf("%-20s %-5s %-10s\n", "Name", "Age", "Gender")
                fmt.Println(strings.Repeat("-", 40))
                for _, person := range people {
                    fmt.Printf("%-20s %-5d %-10s\n", person.Name, person.Age, person.Gender)
                }
            }

        case 3:
            fmt.Println("Exiting CLI...")
            return

        default:
            fmt.Println("Invalid choice, try again.")
        }
    }
}
