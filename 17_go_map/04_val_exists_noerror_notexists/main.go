package main

import "fmt"

func main() {
	g1 := map[int]string{
		1: "Node.js",
		2: "Java",
		3: "Golang",
		4: "Ruby",
		5: "Python",
	}
	// comma ok idiom
	if val, exists := g1[1]; exists {
		delete(g1, 2)
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)
	} else {
		fmt.Println("That value doesn't exists")
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)
	}

	fmt.Println(g1)

	// Check out https://www.golang-book.com/books/intro/6
	// Golang Spec
	// The comparison operators == and != must be fully defined
	// for operands of the key type;
	// thus the key type must not be a
	// function, map, or slice.

}
