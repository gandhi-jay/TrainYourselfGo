package main

import "fmt"

func main() {
	var x [58]string
	// if we define the size it's array
	// otherwise it's slice.
	// array not dynamic. Fixed in Size
	fmt.Println(x)      // Printing whole array
	fmt.Println(len(x)) // length of array
	fmt.Println(x[12])  // 13th element. Array is zero indexed based.
	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}
	fmt.Println(x)      // Printing whole array
	fmt.Println(len(x)) // length of array
	fmt.Println(x[12])  // 13th element. Array is zero indexed based.

	// INT based - default value will be 0s.
	var y [58]int
	// if we define the size it's array
	// otherwise it's slice.
	// array not dynamic. Fixed in Size
	fmt.Println(y)      // Printing whole array
	fmt.Println(len(y)) // length of array
	fmt.Println(y[12])  // 13th element. Array is zero indexed based.
	y[32] = 12321
	fmt.Println(y[32])
}
