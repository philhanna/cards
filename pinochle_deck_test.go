package cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPinochleDeckSort(t *testing.T) {
	pinDeck := NewPinochleDeck(
		Card{TEN, CLUBS},
		Card{JACK, DIAMONDS},
		Card{TEN, SPADES},
	)
	pinDeck.Sort()
	assert.Equal(t, Card{JACK, DIAMONDS}, pinDeck.Cards[0])
	assert.Equal(t, Card{TEN, CLUBS}, pinDeck.Cards[1])
	assert.Equal(t, Card{TEN, SPADES}, pinDeck.Cards[2])
}

func TestPinochleDeck(t *testing.T) {
	deck := NewPinochleDeck()
	PrintDeck("New pinochle deck", deck.Cards)
}

func TestPinochleDeckShuffled(t *testing.T) {
	deck := NewPinochleDeck()
	deck.Shuffle()
	PrintDeck("New pinochle deck shuffled", deck.Cards)
}

func TestPinochleDeckShuffledThenSorted(t *testing.T) {
	deck := NewPinochleDeck()
	deck.Shuffle()
	deck.Sort()
	PrintDeck("New pinochle deck shuffled then sorted", deck.Cards)
}

func TestPinochleRemove(t *testing.T) {
	var ACE_OF_SPADES = NewCard(ACE, SPADES)
	deck := NewPinochleDeck()
	removed := deck.Remove(ACE_OF_SPADES)
	assert.True(t, removed)
	nAceSpades := 0
	for _, card := range deck.Cards {
		if card == ACE_OF_SPADES {
			nAceSpades++
		}
	}
	assert.Equal(t, 47, deck.Len())
	assert.Equal(t, 1, nAceSpades)
}

func TestPinochleSort2(t *testing.T) {
	var (
		TEN_OF_HEARTS = NewCard(TEN, HEARTS)
		KING_OF_CLUBS = NewCard(KING, CLUBS)
	)
	deck := NewPinochleDeck(
		TEN_OF_HEARTS,
		KING_OF_CLUBS,
	)
	deck.Sort()
	assert.Equal(t, deck.Cards[0], KING_OF_CLUBS)
	assert.Equal(t, deck.Cards[1], TEN_OF_HEARTS)
}
