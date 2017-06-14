package main

import (
	"fmt"
	"sort"
)

// Method Sets
// A type may have a method set associated with it.
// The method set of an interface type is its interface.
// The method set of any other type T consists of all methods declared with receiver type T.

func main() {
	s := []string{"WonderWoman", "Batman", "Superman", "Flash", "Arrow"}
	fmt.Println(s)

	fmt.Printf("%T \n", sort.StringSlice(s)) // sort.StringSlice

	// If this is StringSlice, you have some method available.
	// like Len, Less, Swap, Search and Sort

	sort.StringSlice(s).Sort()

	fmt.Println(s)
}
