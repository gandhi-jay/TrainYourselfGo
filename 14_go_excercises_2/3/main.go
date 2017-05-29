package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(greatest(12, 312, 32, 1, 32, 54, 23, 65))
}

func greatest(slice ...int) int {
	sort.Ints(slice)
	return slice[len(slice)-1]
}
