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

	cards := []string{newCard(), "Seven of Spades"}

	cards = append(cards, "Seven of Diamonds")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace of Spades"
}
