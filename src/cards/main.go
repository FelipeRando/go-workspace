package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	// cards.print()
	// fmt.Println("The length of this deck is ", cards.countCards(), "cards")

	// hand, remainingdeck := deal(cards, 3)
	// fmt.Println("Hand:", hand, "\n", "deck:", remainingdeck)

	fmt.Println(cards.toString())
	cards.saveToFile("cartas.txt")
}
