package cards

import "math/rand"
import "slices"

// -----------------------------------------------------------------------
// Type definitions
// -----------------------------------------------------------------------

// Deck is the abstract base type for RegularDeck and PinochleDeck
type Deck []Card

// -----------------------------------------------------------------------
// Methods
// -----------------------------------------------------------------------

// Cards returns the underlying array of cards
func (d *Deck) Cards() []Card {
	return []Card(*d)
}

// Remove removes the specified card from the deck, if it exists.
// Returns true if the card was removed.
func (d *Deck) Remove(c Card) bool {
	p := slices.Index(*d, c)
	if p == -1 {
		return false
	}
	cards := d.Cards()
	cards = append(cards[:p], cards[p+1:]...)
	*d = cards
	return true
}

// Shuffle randomizes the order of cards in a Deck
func (d *Deck) Shuffle() {
	rand.Shuffle(len(*d), func(i int, j int) {
		dc := []Card(*d)
		dc[i], dc[j] = dc[j], dc[i]
	})
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
