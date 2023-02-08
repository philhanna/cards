package cards

import (
	"fmt"
	"strings"
	"testing"
)

func PrintDeck(label string, deck *Deck) {
	fmt.Printf("%s:\n", label)
	ss := make([]string, 0)
	hs := make([]string, 0)
	ds := make([]string, 0)
	cs := make([]string, 0)
	for _, card := range deck.Cards {
		cardString := card.Unicode()
		switch card.Suit {
		case SPADES:
			ss = append(ss, cardString)
		case HEARTS:
			hs = append(hs, cardString)
		case DIAMONDS:
			ds = append(ds, cardString)
		case CLUBS:
			cs = append(cs, cardString)
		}
	}
	spades := strings.Join(ss, ",")
	hearts := strings.Join(hs, ",")
	diamonds := strings.Join(ds, ",")
	clubs := strings.Join(cs, ",")
	fmt.Printf("%s\n%s\n%s\n%s\n", spades, hearts, diamonds, clubs)
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
