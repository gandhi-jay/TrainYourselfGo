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
//
// func Sort
// func Sort(data Interface)
// Sort sorts data.
// It makes one call to data.Len to determine n, and O(n*log(n)) calls to data.Less
// and data.Swap. The sort is not guaranteed to be stable.

type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	dcUniverse := people{"Batman", "Superman", "Flash", "Arrow", "WonderWoman"}
	sort.Sort(dcUniverse)
	fmt.Println(dcUniverse)
}
