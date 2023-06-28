package main

func main() {
	//both of these lines accomplish the same thing
	//var card string = "Ace of Spades"
	// cards := "Ace of Spades"
	//this is a slice, and all data types in the slice must be the same.
	cards := newDeck()
	cards.saveToFile("my_cards.txt")
}