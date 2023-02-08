package cards

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func PrintDeck(label string, cards []Card) {
	fmt.Printf("%s:\n", label)
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
	PrintDeck("New deck", []Card(*deck))
}

func TestDeckShuffled(t *testing.T) {
	deck := NewDeck()
	deck.Shuffle()
	PrintDeck("Shuffled deck", []Card(*deck))
}

func TestDeckShuffledThenSorted(t *testing.T) {
	deck := NewDeck()
	deck.Shuffle()
	sort.Sort(deck)
	PrintDeck("Shuffled then sorted deck", []Card(*deck))
}

func TestPinochleDeck(t *testing.T) {
	deck := NewPinochleDeck()
	PrintDeck("New pinochle deck", []Card(*deck))
}

func TestPinochleDeckShuffled(t *testing.T) {
	deck := NewPinochleDeck()
	deck.Shuffle()
	PrintDeck("New pinochle deck shuffled", []Card(*deck))
}

func TestPinochleDeckShuffledThenSorted(t *testing.T) {
	deck := NewPinochleDeck()
	deck.Shuffle()
	sort.Sort(deck)
	PrintDeck("New pinochle deck shuffled then sorted", []Card(*deck))
}
