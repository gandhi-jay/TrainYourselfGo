package main

import "fmt"

func main() {
	// slice is zero indexed array
	// Slice vs. Slicing vs. index access
	mySlice := []int{1, 3, 4, 56, 7, 23, 56, 32, 31}
	fmt.Println(mySlice, "Length of Slice", len(mySlice))
	// Slicing a slice
	fmt.Println(mySlice[2:6])  // Elements 2 to 5, excluding 6th index
	fmt.Println(mySlice[2])    // 2nd element
	fmt.Println("myString"[2]) // S ascii.
	/*
		You can do "slicing" on string because a string is slice of bytes
		Remember:
			A string is made of runes
			A rune is unicode code point
			A unicode code point is 1 to 4 bytes.
		Strings are made up of runes, runes are made of bytes so string is made up of bytes
		A String is a bunch of bytes; a slice of bytes
	*/
	// Checking capacity of slice
	fmt.Println("Capacity of slice ", cap(mySlice))
}
