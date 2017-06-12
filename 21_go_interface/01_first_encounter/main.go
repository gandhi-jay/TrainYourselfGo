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

// Another Shape
type circle struct {
	radius float64
}

func (z circle) area() float64 {
	return math.Pi * z.radius * z.radius
}

// interface defines functionality
// method should have same signature
// We have same method signature for square
// so we can say,
// 	square implements shape interface
// 	because of that you can create func which takes shape as arg
// signature have to match
type shape interface {
	area() float64
}

// Method takes any type which implements shape
func info(z shape) {
	fmt.Println(z)        // {10} / {10}
	fmt.Printf("%T\n", z) // main.square / main.circle
	fmt.Println(z.area()) // 100 / 314.159..
}

func main() {
	s := square{10}
	// Still i can call s.area()
	info(s) // Calling info on square

	c := circle{10}
	// Still i can call c.area()
	info(c) // calling info on circle
}

// Polymorphism
