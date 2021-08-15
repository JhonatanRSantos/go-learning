package main

import (
	"github.com/JhonatanRSantos/go-learning/V1/deck"
)

func main() {
	cards := deck.NewDeck()
	cards.Shuffle()
	cards.Print()
}
