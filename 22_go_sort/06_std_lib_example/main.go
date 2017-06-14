package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int
}

// Like java's toString method.
// Formatted printing in Go uses a style similar to C's printf family but is richer and more general.
// The functions live in the fmt package and have capitalized names: fmt.Printf, fmt.Fprintf, fmt.Sprintf and so on.
// The string functions (Sprintf etc.) return a string rather than filling in a provided buffer.
func (p person) String() string {
	return fmt.Sprintf("%s, %d\n", p.name, p.age)
}

// To sort by age
type byAge []person

func (a byAge) Len() int {
	return len(a)
}

func (a byAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byAge) Less(i, j int) bool {
	return a[i].age < a[j].age
}

func main() {
	people := []person{
		{"Berry", 31},
		{"Bruce", 42},
		{"Clark", 47},
		{"Arthur", 26},
	}

	fmt.Println(people[0])

	fmt.Printf("%T\n", byAge(people)) // main.byAge
	// Sorting people
	sort.Sort(byAge(people))
	fmt.Println(people)
}

// https://golang.org/pkg/sort/#Sort
// https://golang.org/pkg/sort/#Interface

// String() string
// https://golang.org/doc/effective_go.html#printing
