package main

import "fmt"

func main() {
	fmt.Println("Hello World, First GO!!!")
	fmt.Println(21)
	fmt.Printf("%d - %b \n", 52, 52)
	fmt.Printf("%d - %b - %x \n", 79, 79, 79)
	fmt.Printf("%#X \n", 79)
	for i := 65; i < 125; i++ {
		fmt.Printf("%d \t %b \t %#X \t %q \n", i, i, i, i)
	}
}
