package main

import "fmt"

// create a new type of 'deck'
// which is a slice of strings

// deck type extends the functionality of a slice of strings
type deck []string // type deck === []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamnonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// loop through to create a deck of cards
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
