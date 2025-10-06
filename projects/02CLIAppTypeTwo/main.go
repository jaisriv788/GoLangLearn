package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    fmt.Println("Addition Calculator (Press Ctrl+C to quit)")

    // Run a goroutine for input
    go func() {
        for {
            var num1, num2 int
            fmt.Print("Enter first number: ")
            fmt.Scanln(&num1)
            fmt.Print("Enter second number: ")
            fmt.Scanln(&num2)
            fmt.Printf("Result: %d + %d = %d\n", num1, num2, num1+num2)
        }
    }()

    // Block until Ctrl+C
	fmt.Println("Server running. Press Ctrl+C to stop.")
    sig := make(chan os.Signal, 1)
    signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
    <-sig
    fmt.Println("Shutting down gracefully...")
    fmt.Println("\nCalculator stopped.")
}
