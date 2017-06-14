package main

import (
	"fmt"
	"sort"
)

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
