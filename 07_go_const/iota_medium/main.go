package main

import (
	"fmt"
)

// Weekday exported
type Weekday int

// Constants without iota
// const (
// 	Sunday    Weekday = 0
// 	Monday    Weekday = 1
// 	Tuesday   Weekday = 2
// 	Wednesday Weekday = 3
// 	Thursday  Weekday = 4
// 	Friday    Weekday = 5
// 	Saturday  Weekday = 6
// )

// Constant with iota
const (
	Sunday Weekday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// iota reset: it will be 0.
const (
	ZeroIota = iota // Zero = 0
	OneIota         // One = 1
)

// Timezone int
type Timezone int

// Const with skip iota
const (
	// iota: 0, EST: -5
	EST Timezone = -(5 + iota)
	// _ is the blank identifier
	// iota: 1
	_
	// iota: 2, MST: -7
	MST
	// iota: 3, MST: -8
	PST
)

// Using iota in the middle
const (
	One = 1
	Two = 2
	// Three = 2 + 1 => 3
	// iota in the middle
	Three = iota + 1
	// Four  = 3 + 1 => 4
	Four
)

func main() {
	// https://blog.learngoprogramming.com/golang-const-type-enums-iota-bc4befd096d3
	fmt.Println(Saturday) // 6
	fmt.Printf("Which day it is? %s\n", Sunday.toString())
	fmt.Printf("Is Saturday a weekend day? %t\n", Saturday.isWeekend())
}

func (day Weekday) toString() string {
	// declare an array of strings
	// ... operator counts how many
	// items in the array (7)
	names := [...]string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}
	// → `day`: It's one of the
	// values of Weekday constants.
	// If the constant is Sunday,
	// then day is 0.

	// prevent panicking in case of
	// `day` is out of range of Weekday
	if day < Sunday || day > Saturday {
		return "Unknown"
	}

	return names[day]
}

func (day Weekday) isWeekend() bool {
	switch day {
	case Sunday, Saturday:
		return true

	default:
		return false
	}
}
