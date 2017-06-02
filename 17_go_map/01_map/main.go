package main

import "fmt"

func main() {
	// A map maps keys to values. unordered
	// Maps built on hashtable
	// value of unintialized map is nil
	// Map is reference type
	m := make(map[string]int)

	m["k1"] = 21
	m["k2"] = 22
	fmt.Println("map", m)
	fmt.Println("length", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	// The optional second return value when getting a
	// value from a map indicates if the key was present
	// in the map. This can be used to disambiguate
	// between missing keys and keys with zero values
	// like `0` or `""`. Here we didn't need the value
	// itself, so we ignored it with the _blank identifier_
	// `_`.
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// you can declared and initialize a new map in
	// the same line with syntax
	n := map[string]int{"foo": 123, "bar": 2}
	fmt.Println("map: ", n)
}
