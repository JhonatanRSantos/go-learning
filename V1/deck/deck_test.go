package deck

import (
	"os"
	"testing"
)

const deckLength = 52
const testFileName = "_deckTesting"

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	if len(d) != deckLength {
		t.Errorf("Expected deck length of %d, but got %d", deckLength, len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected first card of Ace of Hearts, but got %s", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %s", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove(testFileName)

	deck := NewDeck()
	deck.SaveToFile(testFileName)

	loadedDeck := NewDeckFromFile(testFileName)

	if len(loadedDeck) != deckLength {
		t.Errorf("Expected deck length of %d, but got %d", deckLength, len(loadedDeck))
	}

	os.Remove(testFileName)
}
