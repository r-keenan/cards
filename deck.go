package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// no receiver, because we are calling this to get a new deck and not calling it on an existing deck
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	
	// use '_' to replace the index variable, because you do not need to use it in these loops
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit )
		}
	}

	return cards
}

// (d deck) is the receiver for the function
// any variable of type 'deck' gets access to the printDeck() method
func (d deck) printDeck() {
	for i, card := range d{
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
