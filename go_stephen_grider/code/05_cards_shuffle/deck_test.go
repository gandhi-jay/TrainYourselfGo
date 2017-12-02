package main

import "testing"
import "os"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected length: 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card: Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card: King of Clubs, but got %v", d[len(d)-1])
	}

	if d[37] != "Queen of Hearts" {
		t.Errorf("Expected 38th card: Queen of Hearts, but got %v", d[37])
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()
	hand, remainingDeck := deal(d, 5)

	if len(hand) != 5 {
		t.Errorf("Expected length of hand: 5, but got %v", len(hand))
	}

	if (len(d) - len(hand)) != len(remainingDeck) {
		t.Errorf("Length of remaining should be remaining length(%v), but it's not equal(%v).", (len(d) - len(hand)), len(remainingDeck))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(d) != len(loadedDeck) {
		t.Errorf("Expected length: %v, but got %v", len(d), len(loadedDeck))
	}

	os.Remove("_decktesting")

}
