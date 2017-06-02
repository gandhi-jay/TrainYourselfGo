package main

import "fmt"

func main() {
	g1 := map[int]string{
		1: "Node.js",
		2: "Java",
		3: "Golang",
		4: "Ruby",
		5: "Python",
	}

	for key, val := range g1 {
		fmt.Println(key, " --> ", val)
	}
}
