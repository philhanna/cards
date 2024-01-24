package cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPinochleDeckSort(t *testing.T) {
	var (
		TEN_OF_CLUBS     = Card{TEN, CLUBS}
		JACK_OF_DIAMONDS = Card{JACK, DIAMONDS}
		TEN_OF_SPADES    = Card{TEN, SPADES}
	)
	cards := []Card{
		TEN_OF_CLUBS, JACK_OF_DIAMONDS, TEN_OF_SPADES,
	}
	pinDeck := PinochleDeck{Deck{cards}}
	pinDeck.Sort()
	assert.Equal(t, Card{JACK, DIAMONDS}, pinDeck.cards[0])
	assert.Equal(t, Card{TEN, CLUBS}, pinDeck.cards[1])
	assert.Equal(t, Card{TEN, SPADES}, pinDeck.cards[2])
}

func TestPinochleDeck(t *testing.T) {
	deck := NewPinochleDeck()
	printDeck("New pinochle deck", deck.cards)
}

func TestPinochleDeckShuffled(t *testing.T) {
	deck := NewPinochleDeck()
	deck.Shuffle()
	printDeck("New pinochle deck shuffled", deck.cards)
}

func TestPinochleDeckShuffledThenSorted(t *testing.T) {
	deck := NewPinochleDeck()
	deck.Shuffle()
	deck.Sort()
	printDeck("New pinochle deck shuffled then sorted", deck.cards)
}

func TestPinochleRemove(t *testing.T) {
	var ACE_OF_SPADES = NewCard(ACE, SPADES)
	deck := NewPinochleDeck()
	removed := deck.Remove(ACE_OF_SPADES)
	assert.True(t, removed)
	nAceSpades := 0
	for _, card := range deck.cards {
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
	cards := []Card{TEN_OF_HEARTS, KING_OF_CLUBS}
	pd := PinochleDeck{Deck{cards}}
	pd.Sort()
	assert.Equal(t, pd.cards[0], KING_OF_CLUBS)
	assert.Equal(t, pd.cards[1], TEN_OF_HEARTS)
}
