package main

import "fmt"

var z = 12
var y int

func main() {

	// Short Declaration Var.
	x := 42 + 7
	fmt.Println(x)
	// x := 49
	// no new variables on left side of :=
	// Can't be used := again for x
	x = 45
	fmt.Println(x)

	// Var Keyword
	foo()

}

func foo() {
	// Gives error because x is out of scope.
	// fmt.Println(x)

	// After declaring z out side of main
	fmt.Println(z)
}
