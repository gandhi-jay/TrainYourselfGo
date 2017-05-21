package main

import "fmt"

func main() {
	b := true
	if food := "Coco"; b {
		fmt.Println(food)
	} else {
		fmt.Println("Nothing available")
	}

	// food will be undefined outside if block

	if false {
		fmt.Println("first print statement")
	} else if false {
		fmt.Println("second print statement")
	} else if true {
		fmt.Println("ahahaha print statement")
	} else {
		fmt.Println("third print statement")
	}

}
