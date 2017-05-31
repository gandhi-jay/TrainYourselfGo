package main

import "fmt"

func main() {
	var x [256]byte
	// byte is alias for uint8 - unsigned integer 8 bits
	for i := 0; i < 256; i++ {
		x[i] = byte(i)
	}
	for i, v := range x {
		fmt.Printf("%v\t%T\t%b \n", v, v, v)
		if i > 50 {
			break
		}
	}
}
