package main

import "fmt"

func main() {
	// Function with One param.
	greet("Jay")
	// Function with Two params.
	greet2("Jay", "Gandhi")
	// Function with two params but mentioning only type one type.
	greet3("Jay", "Jay")
}

// Main func is entry point. starting execution point.

func greet(name string) {
	fmt.Println("Hello", name)
}

func greet2(fname string, lname string) {
	fmt.Println("Hello", fname, lname)
}

func greet3(fname, lname string) {
	greet2(fname, lname)
}
