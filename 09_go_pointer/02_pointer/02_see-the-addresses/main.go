package main

import "fmt"

func zero(z *int) {
	fmt.Println("In Method", z)
	*z = 0
}

// Watch Stephen Grider's Pointer.
func main() {
	x := 5
	fmt.Println(&x)
	zero(&x)
	fmt.Println(x) // x is 0
}
