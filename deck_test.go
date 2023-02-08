package cards

import (
	"sort"
	"testing"
)

func TestDeck(t *testing.T) {
	deck := NewDeck()
	printDeck("New deck", []Card(*deck))
}

func TestDeckShuffled(t *testing.T) {
	deck := NewDeck()
	deck.Shuffle()
	printDeck("Shuffled deck", []Card(*deck))
}

func TestDeckShuffledThenSorted(t *testing.T) {
	deck := NewDeck()
	deck.Shuffle()
	sort.Sort(deck)
	printDeck("Shuffled then sorted deck", []Card(*deck))
}
