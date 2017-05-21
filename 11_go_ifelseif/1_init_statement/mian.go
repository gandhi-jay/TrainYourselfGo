package main

import "fmt"

func main() {
	b := true
	if food := "Coco"; b {
		fmt.Println(food)
	} else {
		fmt.Println("Nothing available")
	}
}
