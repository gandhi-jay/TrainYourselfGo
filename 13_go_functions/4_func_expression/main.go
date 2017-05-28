package main

import "fmt"

func main() {

	// Function expression
	// function assigned to variable
	greet := func() {
		fmt.Println("Hello")
	}
	greet()
	fmt.Printf("%T \n", greet) // Type is func()

	// Another way function expression with closure
	d := makeGreeter("Jay")
	fmt.Println(d())
}

func makeGreeter(name string) func() string {
	return func() string {
		return fmt.Sprint("Hello, ", name)
	}
}
