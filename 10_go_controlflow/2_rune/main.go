package main

import "fmt"

func main() {
	// Watch Rune Lecture from Todd McLoad
	fmt.Println([]byte("Jay"), " -- Jay")
	fmt.Println([]byte("જય"), " -- જય")
	fmt.Println([]byte("જ"), " -- જ")
	for i := 50; i < 140; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	// Don't put i in single quote
}
