package cards

import (
	"sort"
	"testing"
)

func TestPinochleDeck(t *testing.T) {
	deck := NewPinochleDeck()
	printDeck("New pinochle deck", []Card(*deck))
}

func TestPinochleDeckShuffled(t *testing.T) {
	deck := NewPinochleDeck()
	deck.Shuffle()
	printDeck("New pinochle deck shuffled", []Card(*deck))
}

func TestPinochleDeckShuffledThenSorted(t *testing.T) {
	deck := NewPinochleDeck()
	deck.Shuffle()
	sort.Sort(deck)
	printDeck("New pinochle deck shuffled then sorted", []Card(*deck))
}
