package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := &person{"Oliver Queen", 23}
	fmt.Println(p1)              // Pointer to struct person
	fmt.Printf("%T\n", p1)       // *main.person
	fmt.Println(p1.name, p1.age) // accessing fields

}
