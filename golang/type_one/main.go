package main

import "fmt"

func main() {
    fmt.Println("Addition Calculator (Type 'exit' to quit)")

    for {
        var input1, input2 string
        fmt.Print("Enter first number: ")
        fmt.Scanln(&input1)
        if input1 == "exit" {
            break
        }

        fmt.Print("Enter second number: ")
        fmt.Scanln(&input2)
        if input2 == "exit" {
            break
        }

        fmt.Printf("Result: %T + %T\n", input1, input2)


        var num1, num2 int
        fmt.Sscan(input1, &num1)
        fmt.Sscan(input2, &num2)
        fmt.Printf("Result: %T + %T\n", num1, num2)

        fmt.Printf("Result: %d + %d = %d\n", num1, num2, num1+num2)
    }

    fmt.Println("Calculator closed.")
}
