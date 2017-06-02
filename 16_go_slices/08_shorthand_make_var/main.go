package main

import "fmt"

func main() {
	// var slice
	var x []int    // Underlying array not made
	var xx [][]int // Underlying array not made
	// Shorthand slice
	y := []string{}    // Underlying array made
	yy := [][]string{} // Underlying array made
	// Make slice
	z := make([]int, 3, 5)    // Underlying array made
	zz := make([][]int, 3, 5) // Underlying array made

	// Can't use, Index out of range
	// y[0] = "Han Solo"
	y = append(y, "Han Solo")

	// Println
	fmt.Println("Printing all - ", x, y, z)
	// Multidimensional
	fmt.Println("Printing all - ", xx, yy, zz)

	// Comparing with nil
	fmt.Println(x == nil) // true, So nothing is created. Header is pointing to nothing, var creates nothing
	fmt.Println(y == nil) // false
	fmt.Println(z == nil) // false

	// Refer for make and new- https://stackoverflow.com/questions/25358130/what-is-the-difference-between-new-and-make
	// and https://dave.cheney.net/2014/08/17/go-has-both-make-and-new-functions-what-gives
}
