package main

func main() {
	//both of these lines accomplish the same thing
	//var card string = "Ace of Spades"
	//this is a slice, and all data types in the slice must be the same.
	cards := newDeck()

	// first return value assigned to hand; second return value assigned to remainingCards
	hand, remainingCards := deal(cards, 5)

	hand.printDeck()
	remainingCards.printDeck()
}