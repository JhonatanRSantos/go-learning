package main

import (
	"github.com/JhonatanRSantos/course/V1/deck"
)

func main() {
	cards := deck.NewDeck()
	cards.Shuffle()
	cards.Print()
}
