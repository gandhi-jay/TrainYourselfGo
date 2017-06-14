package main

import (
	"fmt"
	"sort"
)

// Most common Interfaces are
// Reader and Writer from IO package.
// https://godoc.org/io#Reader
// https://godoc.org/io#Writer
//
// Now sort package,
// It has "Interface" interface (https://godoc.org/sort#Interface)
// to have implement this interface, you have to have
// Len, Less and Swap
//
// An interface type specifies a method set called its interface.
// A variable of interface type can store a value of any type with a method set that is any superset of the interface.
// Such a type is said to implement the interface.
// The value of an uninitialized variable of interface type is nil.
//
// Excersise -
// (1)
// type people []string
// studyGroup := people{"Zeno", "John", "Al", "Jenny"}

func main() {
	type people []string
	dcUniverse := people{"Batman", "Superman", "Flash", "Arrow", "WonderWoman"}
	// Calling sort method on dcUniverse
	// sort.Sort(dcUniverse) // Will give error if Len, Less and Swap not implemented

	// to sort string.
	fmt.Println("Not Sorted", dcUniverse)
	sort.Strings(dcUniverse)
	fmt.Println("After Sort", dcUniverse)
	fmt.Println("Verifying Strings are sorted:", sort.StringsAreSorted(dcUniverse))
}
