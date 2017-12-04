package main

import "fmt"

func main() {
	// "Hello World"

	// This is an untyped string constant,
	// which is to say it is a constant textual value that does not yet have a fixed type.
	// Yes, it's a string, but it's not a Go value of type string.
	// It remains an untyped string constant even when given a name:

	const hello = "Hello World"

	// After this declaration, hello is also an untyped string constant.
	// An untyped constant is just a value,
	// one not yet given a defined type that would force it to obey the strict rules
	// that prevent combining differently typed values.

	// It is this notion of an untyped constant that makes it possible for us
	// to use constants in Go with great freedom.

	const typedHello string = "Hello, 世界" // typed constant.

	fmt.Printf("%T", typedHello) // string

}
