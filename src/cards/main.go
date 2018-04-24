package main



func main() {
	// cards.print()
	// fmt.Println("The length of this deck is ", cards.countCards(), "cards")

	// hand, remainingdeck := deal(cards, 3)
	// fmt.Println("Hand:", hand, "\n", "deck:", remainingdeck)

	cards := newDeckFromFile("cartas.txt")
	cards.shuffle()
	cards.print()

}
