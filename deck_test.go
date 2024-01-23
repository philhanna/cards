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

func TestCompare(t *testing.T) {
	tests := []struct {
		name  string
		iCard Card
		jCard Card
		ranks []Rank
		want  int
	}{
		{
			name:  "Less",
			iCard: NewCard(THREE, SPADES),
			jCard: NewCard(FOUR, DIAMONDS),
			ranks: RegularRanks,
			want:  -1,
		},
		{
			name:  "Equal",
			iCard: NewCard(THREE, SPADES),
			jCard: NewCard(THREE, SPADES),
			ranks: RegularRanks,
			want:  0,
		},
		{
			name:  "Greater",
			iCard: NewCard(FOUR, DIAMONDS),
			jCard: NewCard(THREE, SPADES),
			ranks: RegularRanks,
			want:  1,
		},
		{
			name:  "LessPinochle",
			iCard: NewCard(KING, SPADES),
			jCard: NewCard(TEN, DIAMONDS),
			ranks: PinochleRanks,
			want:  -1,
		},
		{
			name:  "EqualPinochle",
			iCard: NewCard(THREE, SPADES),
			jCard: NewCard(THREE, SPADES),
			ranks: PinochleRanks,
			want:  0,
		},
		{
			name:  "GreaterPinochle",
			iCard: NewCard(TEN, DIAMONDS),
			jCard: NewCard(KING, SPADES),
			ranks: PinochleRanks,
			want:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			have := Compare(tt.iCard, tt.jCard, tt.ranks)
			switch {
			case tt.want < 0:
				assert.Negative(t, have)
			case tt.want == 0:
				assert.Zero(t, have)
			case tt.want > 0:
				assert.Positive(t, have)
			}
		})
	}
}
