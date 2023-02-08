package cards

import "math/rand"

// -----------------------------------------------------------------------
// Type definitions
// -----------------------------------------------------------------------

// A Deck is an ordered list of cards.
type Deck []Card

// Ranks is a slice of Rank objects
type Ranks []Rank

// -----------------------------------------------------------------------
// Constructors
// -----------------------------------------------------------------------

// NewDeck constructs a 52-card regular deck and returns a pointer to it.
func NewDeck() *Deck {
	cards := make(Deck, 0)
	for i := 0; i < 4; i++ {
		suit := Suit(i)
		for j := 14; j >= 2; j-- {
			rank := Rank(j)
			card := NewCard(rank, suit)
			cards = append(cards, card)
		}
	}
	return &cards
}

// NewPinochleDeck constructs a 48-card Pinochle deck and returns a pointer to it.
func NewPinochleDeck() *Deck {
	cards := make(Deck, 0)
	ranks := []Rank{ACE, TEN, KING, QUEEN, JACK, NINE}
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			suit := Suit(j)
			for _, rank := range ranks {
				card := NewCard(rank, suit)
				cards = append(cards, card)
			}
		}
	}
	return &cards
}

// -----------------------------------------------------------------------
// Methods
// -----------------------------------------------------------------------

// Shuffle randomizes the order of cards in a Deck
func (d *Deck) Shuffle() {
	rand.Shuffle(len(*d), func(i int, j int) {
		dc := []Card(*d)
		dc[i], dc[j] = dc[j], dc[i]
	})
}

// -----------------------------------------------------------------------
// Implementation of the sort interface
// -----------------------------------------------------------------------

// Len returns the number of Rank objects in the array.
func (rs Ranks) Len() int {
	return len(rs)
}

// Less compares two Rank objects and returns true if the first
// one is less than the second one.
func (rs Ranks) Less(i int, j int) bool {
	return rs[i] < rs[j]
}

// Exchanges two Rank objects in the array.
func (rs Ranks) Swap(i int, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

// PinochleRanks is a slice of Rank objects in a Pinochle deck
type PinochleRanks []Rank

// Less compares two Rank objects in a Pinochle deck and returns true
// if the first one is less than the second one.
func (prs PinochleRanks) Less(i int, j int) bool {
	var rankI, rankJ int
	rankI = pinochleRank(prs[i])
	rankJ = pinochleRank(prs[j])
	return rankI < rankJ
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
