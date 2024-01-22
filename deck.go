package cards

import "math/rand"

// -----------------------------------------------------------------------
// Type definitions
// -----------------------------------------------------------------------

// Deck is the abstract base type for RegularDeck and PinochleDeck
type Deck []Card

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

// Cards returns the underlying array of cards
func (d *Deck) Cards() []Card {
	return []Card(*d)
}

// -----------------------------------------------------------------------
// Implementation of the sort interface. RegularDeck and PinochleDeck
// have their own Less() methods.
// -----------------------------------------------------------------------

// Len returns the number of cards in the deck.
func (d Deck) Len() int {
	return len(d.Cards())
}

// Exchanges two Rank objects in the array.
func (d Deck) Swap(i int, j int) {
	dc := d.Cards()
	dc[i], dc[j] = dc[j], dc[i]
}
