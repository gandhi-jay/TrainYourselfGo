package main

import "fmt"

var a int

type dosa int

var b dosa

func main() {
	a = 21
	b = 12
	fmt.Println(a)        // 21
	fmt.Printf("%T\n", a) // int
	fmt.Println(b)        // 12
	fmt.Printf("%T\n", b) // main.dosa
	// b's type is dosa
	// a = b it's like int = dosa.
	// different types.
}
