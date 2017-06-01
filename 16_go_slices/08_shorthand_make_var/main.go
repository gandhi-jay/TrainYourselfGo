package main

import "fmt"

func main() {
	// var slice
	var x []int
	var xx [][]int
	// Shorthand slice
	y := []string{}
	yy := [][]string{}
	// Make slice
	z := make([]int, 3, 5)
	zz := make([][]int, 3, 5)

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
