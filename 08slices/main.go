package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	//Declaration

	var arr = []int{1, 2, 3}
	fmt.Printf("Type of the array is %T\n", arr)
	fmt.Println("Array is", arr)
	fmt.Println("Array length is", len(arr))

	arr = append(arr, 4, 5, 6, 7, 8)
	fmt.Println("Array is", arr)
	fmt.Println("Array length is", len(arr))

	//Way to slice an array
	arr = append(arr[1:]) // here the ending limit is the length of the slice
	fmt.Println("Array is after sliceing from start", arr)

	arr = append(arr[:6]) // here the starting limit is the 0
	fmt.Println("Array is after sliceing from end", arr)

	arr = append(arr[2:5]) // here the ending and starting limit are present
	fmt.Println("Array is after sliceing from both end", arr)

	highscore := make([]int, 4)
	highscore[0] = 123
	highscore[1] = 789
	highscore[2] = 456
	highscore[3] = 963

	//if we add a value more than the length assigned by the below method it will crash
	// highscore[4] = 777

	//But if do it like this then it will work perflectly
	highscore = append(highscore, 111, 222, 333)

	fmt.Println("Before sorting", highscore)
	fmt.Println("Is highscore sorted", sort.IntsAreSorted(highscore))

	sort.Ints(highscore)

	fmt.Println("Before sorting", highscore)
	fmt.Println("Is highscore sorted", sort.IntsAreSorted(highscore))

	//How to remove the elements based on the index
	var lang = []string{"React", "Vue", "Angular", "Next", "Gatsby"}
	fmt.Println("Front end frameworks", lang)

	removingIndex := 3

	lang = append(lang[:removingIndex], lang[removingIndex+1:]...)
	fmt.Println("Front end frameworks after rmoving the element", lang)

	fmt.Println("\nPress Enter to exit...")
	fmt.Scanln()
}
