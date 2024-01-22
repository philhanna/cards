package cards

import "slices"

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

// RegularDeck is a subtype of the abstract type Deck
type RegularDeck struct {
	Deck
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewRegularDeck constructs a 52-card regular deck and returns a pointer to it.
func NewRegularDeck() *RegularDeck {
	rd := new(RegularDeck)
	cards := make([]Card, 0)
	for _, suit := range Suits {
		for _, rank := range rd.Ranks() {
			card := NewCard(rank, suit)
			cards = append(cards, card)
		}
	}
	rd.Deck = cards
	return rd
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Less is used in the sort interface
func (rd RegularDeck) Less(i, j int) bool {
	rankIndexI := slices.Index(rd.Ranks(), rd.Cards()[i].Rank)
	rankIndexJ := slices.Index(rd.Ranks(), rd.Cards()[j].Rank)
	return rankIndexI < rankIndexJ
}

// Ranks returns a list of ranks in ascending order for a regular deck.
func (rd RegularDeck) Ranks() []Rank {
	return []Rank{
		TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING, ACE,
	}
}
