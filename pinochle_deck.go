package cards

import "slices"

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

// NewPinochleDeck constructs a 48-card Pinochle deck and returns a pointer to it.
func NewPinochleDeck() *PinochleDeck {
	pd := new(PinochleDeck)
	cards := make([]Card, 0)
	for i := 0; i < 2; i++ {
		for _, suit := range Suits {
			for _, rank := range pd.Ranks() {
				card := NewCard(rank, suit)
				cards = append(cards, card)
			}
		}
	}
	pd.Deck = cards
	return pd
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Less is used in the sort interface
func (pd PinochleDeck) Less(i, j int) bool {
	rankIndexI := slices.Index(pd.Ranks(), pd.Cards()[i].Rank)
	rankIndexJ := slices.Index(pd.Ranks(), pd.Cards()[j].Rank)
	return rankIndexI < rankIndexJ
}

// Ranks returns the ranks in order for a Pinochle deck
func (pd PinochleDeck) Ranks() []Rank {
	return []Rank{
		NINE, JACK, QUEEN, KING, TEN, ACE,
	}
}
