package main

import "fmt"

func main() {
	a := 21
	fmt.Println("A is ", a)
	fmt.Println("A is at ", &a)
	fmt.Printf("%d\n", &a)
}
