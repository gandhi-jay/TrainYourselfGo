package main

import (
	"fmt"
)

func main() {
	cards := newDeck()

	fmt.Println(cards)

	cards.print()
}

func newCard() string {
	return "Ace of Spades"
}
