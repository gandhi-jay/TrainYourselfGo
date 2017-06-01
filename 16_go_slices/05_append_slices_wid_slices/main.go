package main

import "fmt"

func main() {
	// Int
	mySliceInt := []int{1, 2, 3, 4, 5}
	myOtherSliceInt := []int{6, 7, 8, 9}

	// refer variadic argument.
	mySliceInt = append(mySliceInt, myOtherSliceInt...)

	fmt.Println(mySliceInt)

	// String
	mySliceStr := []string{"Monday", "Tuesday"}
	myOtherSliceStr := []string{"Wednesday", "Thursday", "Friday"}

	// refer variadic argument
	mySliceStr = append(mySliceStr, myOtherSliceStr...)

	fmt.Println(mySliceStr)
}
