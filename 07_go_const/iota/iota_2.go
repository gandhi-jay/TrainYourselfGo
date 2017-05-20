package main

import "fmt"

const (
	_  = iota
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB
	GB
)

func main() {
	fmt.Println("Binary\t\tDecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
}
