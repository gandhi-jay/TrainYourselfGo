package main

import "fmt"

func main() {
	fmt.Println(add(123))
	fmt.Println(add(124))
}

func add(n int) (int, bool) {
	return n / 2, n%2 == 0
}
