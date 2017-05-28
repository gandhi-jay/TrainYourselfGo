package main

import "fmt"

func main() {
	func() {
		fmt.Println("NEVER USE THIS IN GOLANG, Not recommanded")
	}()
}
