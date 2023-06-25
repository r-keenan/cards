package main

func main() {
	//both of these lines accomplish the same thing
	//var card string = "Ace of Spades"
	//this is a slice, and all data types in the slice must be the same.
	cards := deck{"Ace of Diamonds", newCard()}	
	cards = append(cards, "Six of Spades")

	cards.printDeck()
}

func newCard() string {
	return "Five of Diamonds"
}