package main

import (
	"fmt"
)

func main() {
	cards := newDeck()

	cards.shuffle()

	fmt.Println(cards)
}
