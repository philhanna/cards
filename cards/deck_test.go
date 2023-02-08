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
		if len(buffer) == 13 {
			fmt.Printf("%s\n", strings.Join(buffer, " "))
			buffer = buffer[0:0]			
		}
	}
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
