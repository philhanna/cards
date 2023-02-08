package cards

import (
	"fmt"
	"strings"
	"testing"
)

func PrintDeck(label string, deck *Deck) {
	fmt.Printf("%s:\n", label)
	buffer := make([]string, 0)
	for _, card := range deck.Cards {
		buffer = append(buffer, card.Unicode())
	}
	fmt.Printf("%s\n", strings.Join(buffer, " "))
}

func TestDeckString(t *testing.T) {
	deck := NewDeck()
	PrintDeck("New deck", deck)
}

func TestDeckShuffled(t *testing.T) {
	deck := NewDeck()
	deck.Shuffle()
	PrintDeck("Shuffled deck", deck)
}
