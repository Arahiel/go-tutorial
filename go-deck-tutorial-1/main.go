package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// hand, _ := deal(cards, 4)

	// hand.saveToFile("myDealCards.txt")
}
