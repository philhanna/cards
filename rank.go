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

// Compares two ranks in a regular deck
func (r Rank) Compare(other Rank) int {
	return compare(int(r), int(other))
}

// Compares two ranks in a Pinochle deck
func (r Rank) ComparePinochle(other Rank) int {
	return compare(int(pinochleRank(r)), int(pinochleRank(other)))
}

// Helper function used by Compare and ComparePinochle
func compare(ir, io int) int {
	switch {
	case ir < io:
		return -1
	case ir > io:
		return 1
	default:
		return 0
	}
}

// Helper function to adjust order for Pinochle decks
func pinochleRank(rank Rank) int {
	switch rank {
	case NINE, JACK, QUEEN, KING:
		return int(rank)
	case TEN:
		return int(KING) + 1
	case ACE:
		return int(KING) + 2
	default:
		return -1
	}
}

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

// Returns the offset in the Unicode system to the beginning of cards
// for this rank
func (r Rank) Offset() int {
	offset := 0
	switch r {
	case TWO:
		offset = 2
	case THREE:
		offset = 3
	case FOUR:
		offset = 4
	case FIVE:
		offset = 5
	case SIX:
		offset = 6
	case SEVEN:
		offset = 7
	case EIGHT:
		offset = 8
	case NINE:
		offset = 9
	case TEN:
		offset = 10
	case JACK:
		offset = 11
	case QUEEN:
		offset = 13
	case KING:
		offset = 14
	case ACE:
		offset = 1
	}
	return offset
}
