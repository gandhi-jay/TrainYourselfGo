package main

import "fmt"

func main() {
	cards := newDeck()

	// fmt.Println(cards)

	// cards.print()

	hand, remainingDeck := deal(cards, 5)

	fmt.Println("Hand")

	hand.print()

	fmt.Println("Remaining Hand")

	remainingDeck.print()

	// Converting string to byte slice.
	// greeting := "Hello, There"
	// fmt.Println([]byte(greeting))

}

func newCard() string {
	return "Ace of Spades"
}
