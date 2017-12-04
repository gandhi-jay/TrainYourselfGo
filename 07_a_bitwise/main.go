package main

import (
	"fmt"
	"math/rand"
)

// &   bitwise AND
// |   bitwise OR
// ^   bitwise XOR
// &^   AND NOT
// <<   left shift
// >>   right shift

func main() {
	// var x uint8 = 0xAC
	// x &= 0xCA
	// fmt.Println(x)
	x := 172
	fmt.Println(x & 1)

	oneToHundred()

	a := 0
	a |= 196
	a |= 456
	fmt.Println(a)
}

func oneToHundred() {
	for x := 0; x < 100; x++ {
		num := rand.Int()
		if num&1 == 1 {
			fmt.Printf("%d is odd\n", num)
		} else {
			fmt.Printf("%d is even\n", num)
		}
	}
}
