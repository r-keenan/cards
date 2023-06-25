package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// (d deck) is the receiver for the function
// any variable of type 'deck' gets access to the printDeck() method
func (d deck) printDeck() {
	for i, card := range d{
		fmt.Println(i, card)
	}
}