package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck
// which is a slice of string.
type deck []string

// Any variable with Type deck in app
// can call this variable.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option 1 - Log the error return a call to newDeck()
		// Option 2 - Log the error and entirely quit the program.
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return strings.Split(string(byteSlice), ",")

}

func (d deck) shuffleDeprecated() {
	for index := range d {
		// Go Random number generation works based on Seed
		// every time seed is same, that's why it's not repeatiting randomize
		// Based on same seed, generator gives same number of sequence.
		newPos := rand.Intn(len(d) - 1)
		d[index], d[newPos] = d[newPos], d[index]
	}
}

func (d deck) shuffle() {
	// Creating new Source for generator based on timestamp
	source := rand.NewSource(time.Now().UnixNano())
	// Giving a generator a new source.
	r := rand.New(source)
	for index := range d {
		// Based on generator on new seed.
		newPos := r.Intn(len(d) - 1)
		d[index], d[newPos] = d[newPos], d[index]
	}
}
