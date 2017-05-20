package main

import "fmt"

func main() {
	var a int = 43
	var b *int = &a
	var c **int = &b
	fmt.Println(&a, *b, **c)
}
