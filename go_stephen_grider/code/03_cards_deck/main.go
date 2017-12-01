package main

import (
	"fmt"
)

func main() {
	cards := deck{newCard(), "Seven of Spades"}

	fmt.Println(cards)

	cards.print()
}

func newCard() string {
	return "Ace of Spades"
}
