package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type doubleO struct {
	person
	first         string
	LicenseToKill bool
}

func main() {
	p1 := doubleO{
		person: person{
			first: "Hunter",
			last:  "Zolomon",
			age:   43,
		},
		first:         "Black Flash",
		LicenseToKill: true,
	}
	fmt.Println(p1)
	p2 := doubleO{
		person: person{
			first: "Barry",
			last:  "Allen",
			age:   -43,
		},
		first:         "Savitar",
		LicenseToKill: true,
	}
	// accessing both parameters
	fmt.Println(p2.first, p2.person.first)
}
