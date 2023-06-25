package main

import "fmt"

func main() {
	//both of these lines accomplish the same thing
	//var card string = "Ace of Spades"
	card := newCard()
	
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}