package main

import (
	"fmt"
)

func main() {
	// var colors map[string]string // Empty Map

	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "ff0000",
		"green": "00ff00",
		"blue":  "0000ff",
		"white": "ffffff",
	}

	colors["white"] = "000000"

	// delete(colors, "red")

	fmt.Println(colors)

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, "~>", hex)
	}
}
