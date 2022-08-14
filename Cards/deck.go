package main

import "fmt"

// Create a new type of "Deck"
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

// usually we refer to the receiver as a 1 or 2
// letters abreviattion
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
