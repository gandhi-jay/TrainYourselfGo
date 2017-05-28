package main

import "fmt"

func main() {

	visit([]int{2, 5, 3, 2}, func(arg2 int) {
		fmt.Println(arg2)
	})

	// Works like Array.map in ES6 - JavaScript
	xs := filter([]int{2, 5, 3, 2}, func(arg2 int) bool {
		return arg2 >= 3
	})
	fmt.Println(xs)

}

// callback - passing function as an argument

func visit(number []int, callback func(int)) {
	for _, v := range number {
		callback(v)
	}
}

func filter(number []int, callback func(int) bool) []int {
	var x []int
	for _, v := range number {
		if callback(v) {
			x = append(x, v)
		}
	}
	return x
}
