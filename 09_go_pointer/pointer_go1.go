package main

import "fmt"

func main() {
	a := 43
	fmt.Println(a, &a)
	var b *int = &a
	fmt.Println(b)
	// Above code makes b a pointer to the memory address where an int is sotred
	// b is of type "int pointer"
	// *int -- the * is part of the type -- b is of type *int
	fmt.Println(*b) // Dereferencing - * is operator
	// b is an int pointer;
	// b points to the memory address where an int is stored
	// to see the value in that memory address, add a * in front of b
	// this is known as dereferencing
	// the * is an operator in this case

}
