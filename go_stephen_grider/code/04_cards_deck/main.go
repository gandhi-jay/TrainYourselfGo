package main

import (
	"fmt"
)

func main() {
	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)

	// fmt.Println(hand.toString())

	// fmt.Println(remainingDeck.toString())

	// cards.saveToFile("my_cards")

	fmt.Println(newDeckFromFile("my_cards"))
}
