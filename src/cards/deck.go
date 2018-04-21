package main

import (
	"fmt"
)

//Create a new type of deck
//which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

	for _, cardSuit := range cardSuits {

		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}

	}

	return cards

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) countCards() int {
	return len(d)
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}
