package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	card := "Ace of Spades"
	// card := "Five of Diamonds" Error
	card = "Five of Diamonds"

	// Calling a method
	card = newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Ace of Spades"
}
