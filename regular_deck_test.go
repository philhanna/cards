package cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRegularDeck(t *testing.T) {
	rd := NewRegularDeck()
	assert.Equal(t, 52, rd.Len())
}

func TestRegularDeckSort(t *testing.T) {
	cards := []Card{
		{TEN, CLUBS},
		{JACK, DIAMONDS},
		{TEN, SPADES},
	}
	rd := RegularDeck{Deck{cards}}
	rd.Sort()
	assert.Equal(t, 3, len(rd.cards))
	assert.Equal(t, Card{TEN, CLUBS}, rd.cards[0])
	assert.Equal(t, Card{TEN, SPADES}, rd.cards[1])
	assert.Equal(t, Card{JACK, DIAMONDS}, rd.cards[2])
}

func TestRegularDeckShuffled(t *testing.T) {
	deck := NewRegularDeck()
	deck.Shuffle()
	printDeck("Shuffled deck", deck.cards)
}

func TestRegularDeckShuffledThenSorted(t *testing.T) {
	deck := NewRegularDeck()
	deck.Shuffle()
	deck.Sort()
	printDeck("Shuffled then sorted deck", deck.cards)
}

func TestRegularDeckRemove(t *testing.T) {
	var ACE_OF_SPADES = NewCard(ACE, SPADES)
	deck := NewRegularDeck()
	removed := deck.Remove(ACE_OF_SPADES)
	assert.True(t, removed)
	nAceSpades := 0
	for _, card := range deck.cards {
		if card == ACE_OF_SPADES {
			nAceSpades++
		}
	}
	assert.Equal(t, 51, deck.Len())
	assert.Equal(t, 0, nAceSpades)
}
