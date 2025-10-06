package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"
)

// backgroundTask runs concurrently and prints status every 10 seconds
func backgroundTask() {
    for {
        fmt.Println("[Background] CLI is running...")
        time.Sleep(10 * time.Second)
    }
}

// calculator handles user input for addition
func calculator() {
    for {
        var num1, num2 int
        fmt.Print("Enter first number: ")
        _, err := fmt.Scanln(&num1)
        if err != nil {
            fmt.Println("Invalid input, try again.")
            continue
        }

        fmt.Print("Enter second number: ")
        _, err = fmt.Scanln(&num2)
        if err != nil {
            fmt.Println("Invalid input, try again.")
            continue
        }

        fmt.Printf("Result: %d + %d = %d\n\n", num1, num2, num1+num2)
    }
}

func main() {
    fmt.Println("=== Addition CLI Calculator ===")
    fmt.Println("Press Ctrl+C to exit\n")

    // Start background task
    // go backgroundTask()

    // Start calculator input handling
    go calculator()

    // Block main thread until Ctrl+C or termination signal
    sig := make(chan os.Signal, 1)
    signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
    <-sig

    fmt.Println("\nCLI application exiting gracefully...")
}
