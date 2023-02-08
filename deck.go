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
	return dc[i].Rank.Compare(dc[j].Rank) < 0
}

// Exchanges two Rank objects in the array.
func (d *Deck) Swap(i int, j int) {
	dc := []Card(*d)
	dc[i], dc[j] = dc[j], dc[i]
}
