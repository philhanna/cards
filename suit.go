package cards

type Suit int

const (
	SPADES Suit = iota
	HEARTS
	DIAMONDS
	CLUBS
)

func (suit Suit) String() string {
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
