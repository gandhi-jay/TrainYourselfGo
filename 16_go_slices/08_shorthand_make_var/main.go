package main

import "fmt"

func main() {
	// var slice
	var x []int
	var xx [][]int
	// Shorthand slice
	y := []int{1, 2, 3}
	yy := [][]int{{1, 2}, {1, 2, 3}}
	// Make slice
	z := make([]int, 3, 5)
	zz := make([][]int, 3, 5)

	// Println
	fmt.Println("Printing all - ", x, y, z)
	// Multidimensional
	fmt.Println("Printing all - ", xx, yy, zz)

	// Comparing with nil
	fmt.Println(x == nil)
	fmt.Println(y == nil)
	fmt.Println(z == nil)
}
