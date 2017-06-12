package main

import "fmt"

type person struct {
	first   string
	last    string
	isAlive bool
}

type speedster struct {
	person
	first string
}

func (p person) Greeting() {
	fmt.Println("Hello", p.first)
}

func (p speedster) Greeting() {
	fmt.Println("Hello", p.first)
}

func main() {
	p1 := person{
		first:   "Eddie",
		last:    "Thawne",
		isAlive: false,
	}
	p2 := speedster{
		person: person{
			first:   "Eobard",
			last:    "Thawne",
			isAlive: true,
		},
		first: "Reverse Flash",
	}

	p1.Greeting()
	p2.Greeting()

	// To access greeting for person in speedster struct
	p2.person.Greeting()
}
