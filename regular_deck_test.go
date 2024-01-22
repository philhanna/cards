package cards

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRegularDeck(t *testing.T) {
	rd := NewRegularDeck()
	assert.Equal(t, 52, rd.Len())
}

func TestRegularDeckSort(t *testing.T) {
	rd := new(RegularDeck)
	cards := []Card{
		{TEN, CLUBS},
		{JACK, DIAMONDS},
		{TEN, SPADES},
	}
	rd.Deck = cards
	sort.Sort(rd)
	assert.Equal(t, 3, len(rd.Cards()))
	assert.Equal(t, Card{TEN, CLUBS}, rd.Cards()[0])
	assert.Equal(t, Card{TEN, SPADES}, rd.Cards()[1])
	assert.Equal(t, Card{JACK, DIAMONDS}, rd.Cards()[2])
}

func TestRegularDeckShuffled(t *testing.T) {
	deck := NewRegularDeck()
	deck.Shuffle()
	PrintDeck("Shuffled deck", deck.Cards())
}

func TestRegularDeckShuffledThenSorted(t *testing.T) {
	deck := NewRegularDeck()
	deck.Shuffle()
	sort.Sort(deck)
	PrintDeck("Shuffled then sorted deck", deck.Cards())
}
