package main

import "fmt"

const meterToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Println("Enter meters Swam: ")
	fmt.Scan(&meters) // Just like C
	yards := meters * meterToYards
	fmt.Println(meters, " meters is ", yards, " yards")
}
