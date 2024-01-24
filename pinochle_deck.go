package cards

// -----------------------------------------------------------------------
// Type definitions
// -----------------------------------------------------------------------

// PinochleDeck is a subtype of Deck for the Pinochle game
type PinochleDeck struct {
	Deck
}

var PinochleRanks = []Rank{
	NINE, JACK, QUEEN, KING, TEN, ACE,
}

// -----------------------------------------------------------------------
// Constructors
// -----------------------------------------------------------------------

// NewPinochleDeck constructs a 48-card Pinochle deck.
func NewPinochleDeck() PinochleDeck {
	cards := make([]Card, 0)
	for i := 0; i < 2; i++ {
		for _, suit := range Suits {
			for _, rank := range PinochleRanks {
				card := NewCard(rank, suit)
				cards = append(cards, card)
			}
		}
	}
	return PinochleDeck{Deck{Cards: cards}}
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Sort sorts the deck in place
func (pd PinochleDeck) Sort() {
	pd.Deck.Sort(PinochleRanks)
}
