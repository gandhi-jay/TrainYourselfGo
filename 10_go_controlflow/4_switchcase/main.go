package main

import (
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "Dipti":
		fmt.Println("Hi Dipti from System")
		fallthrough
	case "Gaurav":
		fmt.Println("Hi Gaurav from System")
		fallthrough
	case "Gandhi":
		fmt.Println("Gandhi is last name")
	default:
		fmt.Println("Hi", os.Args[1], "from User")
	}

	/*
	  no default fallthrough
	  fallthrough is optional
	  -- you can specify fallthrough by explicitly stating it
	  -- break isn't needed like in other languages - Java
	*/
	switch os.Args[2] {
	case "Rushit", "Bhads":
		fmt.Println("Hey Rushit Doshi!!")
	case "Dixit", "Bond":
		fmt.Println("Hey Dixit !!")
	default:
		fmt.Println("System doesn't recognise you")
	}

	myFriendsName := "Ti"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Wassup my friend with name of length 2")
	case myFriendsName == "Tim":
		fmt.Println("Wassup Tim")
	case myFriendsName == "Jenny":
		fmt.Println("Wassup Jenny")
	case myFriendsName == "Marcus", myFriendsName == "Medhi":
		fmt.Println("Your name is either Marcus or Medhi")
	case myFriendsName == "Julian":
		fmt.Println("Wassup Julian")
	case myFriendsName == "Sushant":
		fmt.Println("Wassup Sushant")
	default:
		fmt.Println("nothing matched; this is the default")
	}

	/*
	  expression not needed
	  -- if no expression provided, go checks for the first case that evals to true
	  -- makes the switch operate like if/if else/else
	  cases can be expressions
	*/

}
