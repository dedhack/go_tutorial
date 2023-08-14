package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := newCard() // := is only used when declaring a new variable, infers the type
	// fmt.Println(card)

	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades") // append does not modify the existing slice, it returns a new slice

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	cards := newDeck()

	// use the deal function to create two new decks
	hand, remainingCards := deal(cards, 5)
	fmt.Println("Hand:")
	hand.print()
	fmt.Println("Remaining Cards:")
	remainingCards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
