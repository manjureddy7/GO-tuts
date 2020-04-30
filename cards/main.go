package main

import "fmt"

func main() {

	// Since newDeck(), print(), deal() are in the same package
	// We dont need to do any import functionality

	cards := newDeck()
	cards.print("All cards are")
	hand, remainingCards := deal(cards, 5)

	hand.print("In Hand")
	remainingCards.print("Remaining Cards")
	cards.saveToFile("mycards.txt")
	returnedDecks := newDeckFromFile("mycards.txt")
	fmt.Println("returned values from file", returnedDecks)

	cards.shuffle()
	cards.print("Suffled Cards")
}
