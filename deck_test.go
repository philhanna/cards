package cards

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func printDeck(label string, cards []Card) {
	fmt.Printf("%s:\n", label)
	buffer := make([]string, 0)
	n := len(cards) / 4
	for _, card := range cards {
		buffer = append(buffer, card.Unicode())
		if len(buffer) == n {
			fmt.Printf("%s\n", strings.Join(buffer, " "))
			buffer = buffer[0:0]
		}
	}
}

func TestDeck_Remove(t *testing.T) {
	tests := []struct {
		name    string
		d       Deck
		c       Card
		want    bool
		newDeck Deck
	}{
		{
			name: "Happy path",
			d: Deck{
				[]Card{
					NewCard(ACE, SPADES),
					NewCard(KING, CLUBS),
					NewCard(KING, DIAMONDS),
					NewCard(ACE, CLUBS),
				},
			},
			c:    NewCard(ACE, SPADES),
			want: true,
			newDeck: Deck{
				[]Card{
					NewCard(KING, CLUBS),
					NewCard(KING, DIAMONDS),
					NewCard(ACE, CLUBS),
				},
			},
		},
		{
			name: "Card not in deck",
			d: Deck{
				[]Card{
					NewCard(ACE, SPADES),
					NewCard(KING, CLUBS),
					NewCard(KING, DIAMONDS),
					NewCard(ACE, CLUBS),
				},
			},
			c:    NewCard(ACE, HEARTS),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			startingLength := len(tt.d.cards)
			want := tt.want
			have := tt.d.Remove(tt.c)
			assert.Equal(t, want, have)
			endingLength := len(tt.d.cards)
			if want {
				assert.Equal(t, startingLength, endingLength+1)
			} else {
				assert.Equal(t, startingLength, endingLength)
			}
		})
	}
}

func TestDeck_Sort(t *testing.T) {
	var (
		TWO_OF_SPADES  = NewCard(TWO, SPADES)
		THREE_OF_CLUBS = NewCard(THREE, CLUBS)
	)
	deck := new(Deck)
	deck.cards = []Card{THREE_OF_CLUBS, TWO_OF_SPADES}
	deck.Sort(RegularRanks)
	assert.Equal(t, deck.cards[0], TWO_OF_SPADES)
	assert.Equal(t, deck.cards[1], THREE_OF_CLUBS)
}
