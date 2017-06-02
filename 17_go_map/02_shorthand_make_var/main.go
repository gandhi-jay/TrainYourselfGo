package main

import "fmt"

func main() {
	var g map[string]string
	var g1 = make(map[string]string)
	// optional capacity can be added
	// g2 :=  make(map[string]string, 100)
	// Golang Spec
	// The initial capacity does not bound its size:
	// maps grow to accommodate the number of items stored in them,
	// with the exception of nil maps.
	// A nil map is equivalent to an empty map except that no elements may be added.
	g2 := make(map[string]string)
	g3 := map[string]string{}

	fmt.Println("g is nil", g == nil)   // g is nil true
	fmt.Println("g1 is nil", g1 == nil) // g is nil false
	fmt.Println("g2 is nil", g2 == nil) // g is nil false
	fmt.Println("g3 is nil", g3 == nil) // g is nil false

	// g["Jay"] = "Hello"
	// g["Gandhi"] = "Halo"
	// fmt.Println(g)
	// panic: assignment to entry in nil map
	// There is no append like function for map
	// in slice, append creates underlying Data structure
	// and then appends it. but in maps, it can't add
	// element to nil map.
	// Never use var g map[string]string for maps

	// using g2, you won't get any error.
	g1["Jay"] = "Hello"
	g1["Gandhi"] = "Halo"
	fmt.Println(g1)

	// It will work with g1, g2 and g3 but not with g1.
}
