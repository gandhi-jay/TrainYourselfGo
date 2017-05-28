package main

import "fmt"

// Package level scope
var x int

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
	d := wrapper()
	fmt.Println(d())
	fmt.Println(d())

	incrementA := wrapper()
	incrementB := wrapper()
	fmt.Println("A:", incrementA())
	fmt.Println("A:", incrementA())
	fmt.Println("B:", incrementB())
	fmt.Println("B:", incrementB())
	fmt.Println("B:", incrementB())
}

func increment() int {
	x++
	return x
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope

anonymous function
a function without a name

func expression
assigning a func to a variable
*/

func wrapper() func() int {
	z := 0
	return func() int {
		z++
		return z
	}
}
