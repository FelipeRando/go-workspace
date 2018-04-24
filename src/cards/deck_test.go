package main

import (
	"testing"
	"os"
)

//how to know what to test?
//what do I really care about?
//easy assertions about the code you're writting

//to test a function always create a test with Test in the begining of the name
//and then put the name of the target function like
func TestNewDeck(t *testing.T) { //the "t" is our test handler/manipulator
	d := newDeck()

	if len(d) != 40 {
		t.Errorf("Expected deck lenght of 20, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d) - 1] != "Ten of Clubs" {
		t.Errorf("Expected last card of Ten of Clubs, but got %v", d[len(d) - 1])
	}


}
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	err := deck.saveToFile("_decktesting")
	if err != nil {
		t.Errorf("Error: %v",err)
		os.Exit(1)
	}

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 40 {
		t.Errorf("Expected 40 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}


