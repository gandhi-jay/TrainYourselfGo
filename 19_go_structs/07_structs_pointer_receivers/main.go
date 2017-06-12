package main

import "fmt"

type muteable struct {
	a int
	b int
}

func (m muteable) stayTheSame() {
	m.a = 5
	m.b = 7
	fmt.Println("Value are changed in method stayTheSame", m)
}

func (m *muteable) mutate() {
	m.a = 5
	m.b = 7
}

func main() {
	m := muteable{a: 53, b: 35}
	fmt.Println("Original Values", m)
	m.stayTheSame()
	fmt.Println("Values didn't change using Non pointer receivers", m)
	fmt.Println("Calling mutate method")
	m.mutate()
	fmt.Println("Values changed using pointer receivers", m)
}

/*
Link : https://goo.gl/LS5QWt
Reasons why you would want to pass by reference as opposed to by value:
1. You want to actually modify the receiver (“read/write” as opposed to just “read”)
2. The struct is very large and a deep copy is expensive
3. Consistency: if some of the methods on the struct have pointer receivers,
		the rest should too. This allows predictability of behavior


With method receivers that take pointers,
Go conveniently allows both pointers and non-pointers
to be passed and it automatically does the conversion

http://openmymind.net/Things-I-Wish-Someone-Had-Told-Me-About-Go/
*/
