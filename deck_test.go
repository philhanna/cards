package cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			d: Deck([]Card{
				NewCard(ACE, SPADES),
				NewCard(KING, CLUBS),
				NewCard(KING, DIAMONDS),
				NewCard(ACE, CLUBS),
			}),
			c:    NewCard(ACE, SPADES),
			want: true,
			newDeck: Deck([]Card{
				NewCard(KING, CLUBS),
				NewCard(KING, DIAMONDS),
				NewCard(ACE, CLUBS),
			}),
		},
		{
			name: "Card not in deck",
			d: Deck([]Card{
				NewCard(ACE, SPADES),
				NewCard(KING, CLUBS),
				NewCard(KING, DIAMONDS),
				NewCard(ACE, CLUBS),
			}),
			c:    NewCard(ACE, HEARTS),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			startingLength := tt.d.Len()
			want := tt.want
			have := tt.d.Remove(tt.c)
			assert.Equal(t, want, have)
			endingLength := tt.d.Len()
			if want {
				assert.Equal(t, startingLength, endingLength+1)
			} else {
				assert.Equal(t, startingLength, endingLength)
			}
		})
	}
}
