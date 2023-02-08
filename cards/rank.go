package cards

import "fmt"

// -----------------------------------------------------------------------
// Type definitions
// -----------------------------------------------------------------------

// A rank is one of the 13 levels of cards in a deck.
type Rank int

const (
	TWO Rank = iota + 2
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
	ACE	
)

// -----------------------------------------------------------------------
// Methods
// -----------------------------------------------------------------------

// Returns a representation of the type as a string
func (r Rank) String() string {
	switch r {
	case TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN:
		return fmt.Sprint(int(r))
	case JACK:
		return "J"
	case QUEEN:
		return "Q"
	case KING:
		return "K"
	case ACE:
		return "A"
	default:
		return "?"
	}
}