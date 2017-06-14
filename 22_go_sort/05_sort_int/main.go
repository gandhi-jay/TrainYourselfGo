package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{3, 2, 34, 23, 54, 213, 452, 32}
	fmt.Println(s)
	sort.Sort(sort.IntSlice(s))
	fmt.Println(s)
}
