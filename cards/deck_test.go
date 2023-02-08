package cards

import (
	"fmt"
	// "sort"
	"strings"
	"testing"
)

func PrintDeck(label string, deck *Deck) {
	fmt.Printf("%s:\n", label)
	cards := []Card(*deck)
	buffer := make([]string, 0)
	n := len(cards) / 4
	for _, card := range cards {
		buffer = append(buffer, card.Unicode())
		if len(buffer) == n {
			fmt.Printf("%s\n", strings.Join(buffer, " "))
			buffer = buffer[0:0]
		}
	}
}

func TestDeck(t *testing.T) {
	deck := NewDeck()
	PrintDeck("New deck", deck)
}

func TestDeckShuffled(t *testing.T) {
	deck := NewDeck()
	deck.Shuffle()
	PrintDeck("Shuffled deck", deck)
}

/*func TestDeckShuffledThenSorted(t *testing.T) {
	deck := NewDeck()
	deck.Shuffle()
	sort.Sort(deck)
	PrintDeck("Shuffled deck", deck)
}*/

func TestPinochleDeck(t *testing.T) {
	deck := NewPinochleDeck()
	PrintDeck("New pinochle deck", deck)
}

func TestPinochleDeckShuffled(t *testing.T) {
	deck := NewPinochleDeck()
	deck.Shuffle()
	PrintDeck("New pinochle deck", deck)
}

/*func TestPinochleDeckShuffledThenSorted(t *testing.T) {
	deck := NewPinochleDeck()
	deck.Shuffle()
	//sort.Sort(deck)
	PrintDeck("New pinochle deck", deck)
}*/
