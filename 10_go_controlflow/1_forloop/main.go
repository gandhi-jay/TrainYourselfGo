package main

import "fmt"

func main() {
	count := 10
	for i := 1; i <= count; i++ {
		fmt.Println(i)
	}
	// Go to effective go for for loop doc.
	//
	// Like C While
	for count > 0 {
		fmt.Println(count)
		count--
	}
	count = 3
	// Infinite loop
	// for {}
	//
	// Nested Loop
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			fmt.Println(i, " --- ", j)
		}
	}
	count = 125
	for {
		if count%5 == 0 {
			fmt.Println("Count - ", count)
		}
		if count < 10 {
			break
		}
		count--
	}
}
