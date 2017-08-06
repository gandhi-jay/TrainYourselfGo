package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello World")
	// n is number of bytes written and err is any err is encountered.
	fmt.Println(n, err)

	// If you don't want to catch an error, throw it in void, use _
	n2, _ := fmt.Println("Code Pollution")
	fmt.Println(n2)
}
