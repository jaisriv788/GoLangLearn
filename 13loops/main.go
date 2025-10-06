package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops.")
	days := []string{}

	days = append(days, "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday")
	fmt.Println(days)

	//Below is the normal for loop usage
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	//This is also verson of for loop but similar to foreach
	for i, v := range days {
		fmt.Printf("Day is %v and Day name is %v\n", i+1, v)
	}

	//Below the for loop similar to the while loop as it starts with condition check
	value := 1
	for value < 10 {
		if value == 7 {
			break
		} else if value == 2 {
			value++
			continue
		} else if value == 5 {
			goto msg
		}
		fmt.Println(value)
		value++
	}

	//The msg lable con have anything loops, conditionals, print statement etc...
msg:
	fmt.Println("You came here by goto.")
}
