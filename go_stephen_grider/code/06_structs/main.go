package main

import (
	"fmt"
)

type contactInfo struct {
	email string
}

type cell struct {
	network string
	area    int
}

type person struct {
	firstName string
	lastName  string
	phone     cell
	contactInfo
}

func main() {
	// Not recommanded because anybody can switch order of order.
	// babaYaga := person{"Jhon", "Wick"}
	// fmt.Println(babaYaga)

	// More explicit
	babaYaga := person{
		firstName: "Jhon",
		lastName:  "Wick",
		phone: cell{
			network: "Seek and Destroy",
			area:    0,
		},
		contactInfo: contactInfo{
			email: "jhon@killboogeyman.com",
		},
	}

	fmt.Println("FirstName: ", babaYaga.firstName)
	fmt.Println("LastName: ", babaYaga.lastName)

	var nilPerson person

	fmt.Printf("%+v\n", nilPerson)

	var han person

	han.firstName = "Han"
	han.lastName = "Solo"
	han.contactInfo.email = "hansolo@repulic.com"
	han.phone.area = 0
	han.phone.network = "Republic"

	fmt.Printf("%+v", han)

}
