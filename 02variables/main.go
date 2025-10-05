package main

import "fmt"

const GlobalVariable int = 400

//  GlobalVariable := 400 (Can not declare variables like this.)

func main() {
	var username string = "Jai"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.92759274082556456
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values
	var usernameAssignment string
	fmt.Println(usernameAssignment)
	fmt.Printf("Assigned Variable is of type: %T \n", usernameAssignment)

	var isLoggedInAssignment bool
	fmt.Println(isLoggedInAssignment)
	fmt.Printf("Assigned Variable is of type: %T \n", isLoggedInAssignment)

	var smallValAssignment uint8
	fmt.Println(smallValAssignment)
	fmt.Printf("Assigned Variable is of type: %T \n", smallValAssignment)

	var smallFloatAssignment float32
	fmt.Println(smallFloatAssignment)
	fmt.Printf("Assigned Variable is of type: %T \n", smallFloatAssignment)

	//implecit way to declare the variable
	var autoAssignVar = "I am assigned automatically"
	fmt.Println(autoAssignVar)
	fmt.Printf("Implecit Variable is of type: %T \n", autoAssignVar)

	//Short hand way to declare variable
	shortHandVarAssignment := 32
	fmt.Println(shortHandVarAssignment)
	fmt.Printf("Shorthand Variable is of type: %T \n", shortHandVarAssignment)

	//useage of global variable
	fmt.Println(GlobalVariable)
	fmt.Printf("Global Variable is of type: %T \n", GlobalVariable)
}
