package main

import "fmt"

func main() {
	// Every Thing in golang is pass by value
	// For string also you can try same thing.
	age := 23
	fmt.Println("Age is", age)
	fmt.Println("Age's address is", &age)

	changeme_nopointer(age)
	fmt.Println("After calling changeme_nopointer method with ", age)
	// Using pointer
	changeme(&age)
	fmt.Println("Value change by function ", age)
}

func changeme_nopointer(z int) {
	fmt.Println("before Value in function changeme_nopointer", z)
	z = 44
	fmt.Println("after Value in function changeme_nopointer", z)
}

func changeme(z *int) {
	fmt.Println("------- Using pointer method")
	fmt.Println("Value passed by to function", z)
	fmt.Println("actual value at memory address", *z)
	*z = 32
	fmt.Println("Memory address not changed", z)
	fmt.Println("value is going to be ", *z)
}
