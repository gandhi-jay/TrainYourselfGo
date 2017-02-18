package main

import "fmt"

func main() {
	increment := wrapper()
	fmt.Println(increment)
	fmt.Println(increment())
	fmt.Println(increment())
}

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}