package main

import "fmt"

func main() {

	// Length
	g := make(map[string]string)
	g["Go"] = "Google"
	g["Java"] = "Oracle"
	g["Video"] = "Netflix"

	fmt.Println("Length of g", len(g))

	// Update & Delete
	g1 := map[int]string{
		1: "Node.js",
		2: "Java",
		3: "Golang",
		4: "Ruby",
		5: "Python",
	}

	fmt.Println(g1) // map[4:Ruby 5:Python 1:Node.js 2:Java 3:Golang]
	g1[4] = "Scala" // Updating
	fmt.Println(g1) // map[1:Node.js 2:Java 3:Golang 4:Scala 5:Python]
	delete(g1, 4)   // Deleting
	fmt.Println(g1) // map[5:Python 1:Node.js 2:Java 3:Golang]

}
