package main

import "fmt"

// Creating Person type
type person struct {
	first string
	last  string
	age   int
}

// You can create your own type
type age int

func main() {
	fmt.Println(person{"Bruce", "Wayne", 28})
	fmt.Println(person{"Barry", "Allen", 21})
	super := person{"Clark", "Kent", 31}
	fmt.Printf("%T\n", super)                       // Type of var super
	fmt.Println(super.first, super.last, super.age) // access field of the struct

	// Checking on your own type.
	var myAge age
	myAge = 21
	fmt.Printf("%T %v \n", myAge, myAge) // Check type and it's value
}
