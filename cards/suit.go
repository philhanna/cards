package cards

import "runtime"

type Suit int

const (
	SPADES Suit = iota
	HEARTS
	DIAMONDS
	CLUBS
)

// Returns the Unicode glyph that represents this suit, i. e. ♠, ♡, ♢, ♣
func (suit Suit) Glyph() string {
	switch suit {
	case SPADES:
		return string('\u2660')
	case HEARTS:
		return string('\u2661')
	case DIAMONDS:
		return string('\u2662')
	case CLUBS:
		return string('\u2663')
	default:
		return "?"
	}
}

// Returns the character that represents this suit, i. e. S, H, D, C
func (suit Suit) Character() string {
	switch suit {
	case SPADES:
		return "S"
	case HEARTS:
		return "H"
	case DIAMONDS:
		return "D"
	case CLUBS:
		return "C"
	default:
		return "?"
	}
}

// Returns the suit as a string
func (suit Suit) String() string {
	if runtime.GOOS == "windows" {
		return suit.Character()
	} else {
		return suit.Glyph()	
	}
}
