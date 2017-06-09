package main

import "fmt"

/*
	A struct is a sequence of named elements, called fields, each of which has a name and a type.
	Field names may be specified explicitly (IdentifierList) or implicitly (AnonymousField).
	Within a struct, non-blank field names must be unique.
	A field declared with a type but no explicit field name is an anonymous field,
	also called an embedded field or an embedding of the type in the struct.
	An embedded type must be specified as a type name T or as a pointer to a non-interface type name *T, and T
	itself may not be a pointer type. The unqualified type name acts as the field name.
*/

// Creating Person type
// Data we call as fields
type person struct {
	first string
	last  string
	age   int
}

// You can create your own type
// Alias type - Not idiomatic
type age int

func main() {
	fmt.Println(person{"Bruce", "Wayne", 28})
	fmt.Println(person{"Barry", "Allen", 21})
	super := person{"Clark", "Kent", 31}
	fmt.Printf("%T\n", super)                       // Type of var super
	fmt.Println(super.first, super.last, super.age) // access field of the struct

	p1 := person{}
	fmt.Println(p1) // It will print dafault values of each field type
	// Checking on your own type.
	var myAge age
	myAge = 21
	fmt.Printf("%T %v \n", myAge, myAge) // Check type and it's value

	// Creating another one
	var yourAge age
	yourAge = 32
	fmt.Printf("%T %v \n", yourAge, yourAge)
	// this won't work:
	//	 fmt.Println(myAge + yourAge)

	// this conversion works:
	fmt.Println(int(myAge) + int(yourAge))
	// User defined types - we declare a new type, age
	// the underlying type of age: int
	//
	// conversion: int(myAge)
	// coverting type age to type int
	//
	// This is an example only. IT is a bad practice to alias types
	// Only do use this kind of aliasing when you require to attach method to a type
	//  see the time package for an example of this
	//   godoc.org/time
	//   type Duration int64
	// Duration has methods attached to it
}
