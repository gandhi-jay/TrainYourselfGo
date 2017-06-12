package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

func (z square) area() float64 {
	return z.side * z.side
}

func (z square) info() {
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(z.area())
}

// Another Shape
type circle struct {
	radius float64
}

func (z circle) area() float64 {
	return math.Pi * z.radius * z.radius
}

// Same functionality second time
func (z circle) info() {
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(z.area())
}

func main() {
	s := square{10}
	// Still i can call s.area()
	s.info() // Calling info of square

	c := circle{10}
	// Still i can call c.area()
	c.info() // calling info of circle

	// Info is tightly coupled with circle and square seperatley.
}

// No Polymorphism
