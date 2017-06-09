package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type doubleO struct {
	person
	LicenseToKill bool
}

func main() {
	p1 := doubleO{
		person: person{
			first: "James",
			last:  "Bond",
			age:   23,
		},
		LicenseToKill: true,
	}
	p2 := doubleO{
		person: person{
			first: "Kevin",
			last:  "Peterson",
			age:   30,
		},
		LicenseToKill: false,
	}

	fmt.Println(p1)
	fmt.Println(p1.first, p1.last, p1.age, p1.LicenseToKill)
	fmt.Println(p2)
	fmt.Println(p2.first, p2.last, p2.age, p2.LicenseToKill)
}
