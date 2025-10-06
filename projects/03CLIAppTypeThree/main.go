package main

import (
    "fmt"
    "time"
)

func backgroundTask() {
    for {
        fmt.Println("Calculator is ready...")
        time.Sleep(10 * time.Second) // periodic message
    }
}

func main() {
    go backgroundTask() // run background task

    for {
        var num1, num2 int
        fmt.Print("Enter first number: ")
        fmt.Scanln(&num1)
        fmt.Print("Enter second number: ")
        fmt.Scanln(&num2)
        fmt.Printf("Result: %d + %d = %d\n", num1, num2, num1+num2)
    }
}
