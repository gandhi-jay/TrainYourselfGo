package main

import "fmt"

func main() {
	fmt.Println("Variadic Params", average(1.3123, 13221.132, 132.213))
	// Demo of Variadic arguments
	data := []float64{2.0, 3.0} // Slice of float64
	n := average(data...)       // Passing Slice/List into function which accepts variadic arg
	fmt.Println("Variadic Args", n)
	// Passing Slice as argument instead of Variadic argument
	fmt.Println("Slice as Arg: ", avg(data))
}

// The final incoming parameter in a function signature may have a type prefixed with ....
// A function with such a parameter is called variadic and
// may be invoked with zero or more arguments for that parameter.
//
// Variadic Params
func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf) // Slice of float64
	total := 0.0
	// Iterate over variadic argument
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

// Actual Slice as parameter
func avg(sf []float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
