package cards

// ---------------------------------------------------------------------
// Type Definitions
// ---------------------------------------------------------------------

// RegularDeck is a subtype of the abstract type Deck
type RegularDeck struct {
	Deck
}

var RegularRanks = []Rank{
	TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING, ACE,
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewRegularDeck constructs a 52-card regular deck.
func NewRegularDeck() RegularDeck {
	cards := make([]Card, 0)
	for _, suit := range Suits {
		for _, rank := range RegularRanks {
			card := NewCard(rank, suit)
			cards = append(cards, card)
		}
	}
	return RegularDeck{Deck{cards}}
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Sort sorts the deck in place using the regular ordering.
func (rd *RegularDeck) Sort() {
	rd.Deck.Sort(RegularRanks)
}
