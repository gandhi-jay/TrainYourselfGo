package main

import "fmt"

func main() {
	for idea := 0; idea < 100; idea++ {
		if idea%2 == 0 {
			fmt.Println("Even ", idea)
		}
	}
}
