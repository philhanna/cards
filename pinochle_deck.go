package cards

import (
	"log"
	"slices"
)

// -----------------------------------------------------------------------
// Type definitions
// -----------------------------------------------------------------------

// PinochleDeck is a subtype of Deck for the Pinochle game
type PinochleDeck struct {
	Deck
}

var (
	PinochleRanks = []Rank{
		NINE, JACK, QUEEN, KING, TEN, ACE,
	}
)

// -----------------------------------------------------------------------
// Constructors
// -----------------------------------------------------------------------

// NewPinochleDeck constructs a Pinochle deck. If cards are specified in
// the parameter list, the deck is created with those cards.  Otherwise,
// a full 48-card deck is created.
func NewPinochleDeck(cards ...Card) PinochleDeck {
	if len(cards) == 0 {
		for i := 0; i < 2; i++ {
			for _, suit := range Suits {
				for _, rank := range PinochleRanks {
					card := NewCard(rank, suit)
					cards = append(cards, card)
				}
			}
		}
	}
	for _, card := range cards {
		if !slices.Contains(PinochleRanks, card.Rank) {
			log.Fatalf("%s is not a valid card in a Pinochle deck", card)
		}
	}
	return PinochleDeck{Deck: NewDeck(cards...)}
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Sort sorts the deck in place
func (pd PinochleDeck) Sort() {
	pd.Deck.Sort(PinochleRanks)
}
