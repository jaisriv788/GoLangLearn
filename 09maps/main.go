package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps")

	lang := make(map[string]string)

	lang["JS"] = "JavaScript"
	lang["TS"] = "TypeScript"
	lang["GO"] = "GoLang"

	fmt.Println("All languages", lang)

	//Below is how you delete an element
	delete(lang, "TS")
	fmt.Println("All languages", lang)

	//Below is how we loop through maps
	for key, value := range lang {
		fmt.Printf("For the key %v, value is %v\n", key, value)
	}
}
