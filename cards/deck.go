package cards

import "math/rand"

// -----------------------------------------------------------------------
// Type definitions
// -----------------------------------------------------------------------

// A Deck is an ordered list of cards.
type Deck []Card

// Ranks is a slice of Rank objects
type Ranks []Rank

// A PinochleDeck is an ordered list of cards.
type PinochleDeck []Card

// PinochleRanks is a slice of Rank objects in a Pinochle deck
type PinochleRanks []Rank

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
func NewPinochleDeck() *PinochleDeck {
	cards := make(PinochleDeck, 0)
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

// Shuffle randomizes the order of cards in a Deck
func (d *PinochleDeck) Shuffle() {
	rand.Shuffle(len(*d), func(i int, j int) {
		dc := []Card(*d)
		dc[i], dc[j] = dc[j], dc[i]
	})
}

// -----------------------------------------------------------------------
// Implementation of the sort interface for regular deck
// -----------------------------------------------------------------------

// Len returns the number of cards in the deck.
func (d *Deck) Len() int {
	return len([]Card(*d))
}

// Less compares two cards and returns true if the first
// one is less than the second one.
func (d *Deck) Less(i int, j int) bool {
	dc := []Card(*d)
	return dc[i].Rank < dc[j].Rank
}

// Exchanges two Rank objects in the array.
func (d *Deck) Swap(i int, j int) {
	dc := []Card(*d)
	dc[i], dc[j] = dc[j], dc[i]
}

// -----------------------------------------------------------------------
// Implementation of the sort interface for pinochle deck
// -----------------------------------------------------------------------

// Len returns the number of cards in the deck.
func (d *PinochleDeck) Len() int {
	return len([]Card(*d))
}
// Less compares two Rank objects in a Pinochle deck and returns true
// if the first one is less than the second one.
func (d *PinochleDeck) Less(i int, j int) bool {
	var rankI, rankJ int
	dc := []Card(*d)
	rankI = pinochleRank(dc[i].Rank)
	rankJ = pinochleRank(dc[j].Rank)
	return rankI < rankJ
}

// Exchanges two Rank objects in the array.
func (d *PinochleDeck) Swap(i int, j int) {
	dc := []Card(*d)
	dc[i], dc[j] = dc[j], dc[i]
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
