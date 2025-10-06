package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

// Person struct to store user data
type Person struct {
    Name   string `json:"name"`
    Age    int    `json:"age"`
    Gender string `json:"gender"`
}

const dataFile = "people.json"

// Load people from JSON file
func loadData() ([]Person, error) {
    var people []Person
    if _, err := os.Stat(dataFile); os.IsNotExist(err) {
        return people, nil // file doesn't exist yet
    }

    data, err := ioutil.ReadFile(dataFile)
    if err != nil {
        return nil, err
    }

    err = json.Unmarshal(data, &people)
    return people, err
}

// Save people to JSON file
func saveData(people []Person) error {
    data, err := json.MarshalIndent(people, "", "  ")
    if err != nil {
        return err
    }
    return ioutil.WriteFile(dataFile, data, 0644)
}

// Truncate strings to avoid breaking table
func truncate(s string, max int) string {
    if len(s) > max {
        return s[:max-3] + "..."
    }
    return s
}

// Display table of people
func displayTable(people []Person) {
    if len(people) == 0 {
        fmt.Println("No people stored yet.")
        return
    }

    maxNameLen := 4 // minimum width for "Name"
    for _, p := range people {
        if len(p.Name) > maxNameLen {
            maxNameLen = len(p.Name)
        }
    }

    fmt.Printf("\n%-*s %-5s %-10s\n", maxNameLen, "Name", "Age", "Gender")
    fmt.Println(strings.Repeat("-", maxNameLen+20))

    for _, p := range people {
        fmt.Printf("%-*s %-5d %-10s\n", maxNameLen, truncate(p.Name, maxNameLen), p.Age, p.Gender)
    }
}

func main() {
    people, err := loadData()
    if err != nil {
        fmt.Println("Error loading data:", err)
        return
    }

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
            if err := saveData(people); err != nil {
                fmt.Println("Error saving data:", err)
            } else {
                fmt.Println("Person added successfully!")
            }

        case 2:
            displayTable(people)

        case 3:
            fmt.Println("Exiting CLI...")
            return

        default:
            fmt.Println("Invalid choice, try again.")
        }
    }
}
