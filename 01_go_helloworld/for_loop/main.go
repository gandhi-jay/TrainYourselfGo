package main

import "fmt"

func main() {

	// %#X will display 0xXXX hex. %q represent equivalent utf-8 character
	for i := 0; i <= 100; i++ {
		fmt.Printf(" %d \t %b \t %x \t %#X \t %q \n",i,i,i,i,i)
		// 100     1100100         64      0X64    'd'
	}
}
