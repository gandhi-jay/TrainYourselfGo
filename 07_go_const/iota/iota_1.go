package main

import "fmt"

const (
	A = iota // 0
	B = iota // 1
)

const (
	D = iota //0
	E
	F
)

const (
	_ = iota
	a = iota * 10 // 1 * 10
	b = iota * 10 // 2 * 10
)

func main() {
	fmt.Println(A, B)
	fmt.Println(D, E, F)
	fmt.Println(a, b)
}
