package main

import "fmt"

type deck []string

func newDeck() deck {
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardVaules := []string{"Ace", "Two", "Three", "Four"}

	for i, suit := range cardSuits {
		for j, value := range cardVaules {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
