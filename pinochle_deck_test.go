package cards

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPinochleDeckSort(t *testing.T) {
	pinDeck := new(PinochleDeck)
	cards := []Card{
		{TEN, CLUBS},
		{JACK, DIAMONDS},
		{TEN, SPADES},
	}
	pinDeck.Deck = cards
	sort.Sort(pinDeck)
	assert.Equal(t, Card{JACK, DIAMONDS}, pinDeck.Cards()[0])
	assert.Equal(t, Card{TEN, CLUBS}, pinDeck.Cards()[1])
	assert.Equal(t, Card{TEN, SPADES}, pinDeck.Cards()[2])
}

func TestPinochleDeck(t *testing.T) {
	deck := NewPinochleDeck()
	PrintDeck("New pinochle deck", deck.Cards())
}

func TestPinochleDeckShuffled(t *testing.T) {
	deck := NewPinochleDeck()
	deck.Shuffle()
	PrintDeck("New pinochle deck shuffled", deck.Cards())
}

func TestPinochleDeckShuffledThenSorted(t *testing.T) {
	deck := NewPinochleDeck()
	deck.Shuffle()
	sort.Sort(deck)
	PrintDeck("New pinochle deck shuffled then sorted", deck.Cards())
}
