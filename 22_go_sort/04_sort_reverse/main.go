package main

import (
	"fmt"
	"sort"
)

// To reverse sort
// https://godoc.org/sort#Interface
//

func main() {
	s := []string{"WonderWoman", "Batman", "Superman", "Flash", "Arrow"}
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
}
