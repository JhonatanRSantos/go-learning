package deck

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

var cardsSuits = [4]string{"Hearts", "Diamonds", "Spades", "Clubs"}
var cardsValues = [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

type Deck []string

func NewDeck() Deck {
	cards := Deck{}

	for _, suitValue := range cardsSuits {
		for _, cardValue := range cardsValues {
			cards = append(cards, fmt.Sprintf("%s of %s", cardValue, suitValue))
		}
	}
	return cards
}

func NewDeckFromFile(filename string) Deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return Deck(strings.Split(string(bs), ","))
}

func (d Deck) Print() {
	for index, card := range d {
		fmt.Printf("Index = %d <==>\tCard = %s\n", index, card)
	}
}

func (d Deck) Deal(handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) ToString() string {
	return strings.Join([]string(d), ",")
}

func (d Deck) SaveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.ToString()), 0777)
}

func (d Deck) Shuffle() {
	for index := range d {
		// Creates a random number from 0 to slice_size - 1
		newPosition := rand.New(
			rand.NewSource(
				time.Now().UnixNano(),
			),
		).Intn(len(d) - 1)

		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
