package cards

import (
	"math/rand"
	"slices"
	"sort"
)

// -----------------------------------------------------------------------
// Type definitions
// -----------------------------------------------------------------------

// Deck is the abstract base type for RegularDeck and PinochleDeck
type Deck struct {
	Cards []Card
}

// -----------------------------------------------------------------------
// Methods
// -----------------------------------------------------------------------

// Len returns the number of cards in the deck
func (d Deck) Len() int {
	return len(d.Cards)
}

// Remove removes the specified card from the deck, if it exists.
// Returns true if the card was removed.
func (d *Deck) Remove(c Card) bool {
	p := slices.Index(d.Cards, c)
	if p == -1 {
		return false
	}
	oldCards := d.Cards[:]
	d.Cards = append(oldCards[:p], oldCards[p+1:]...)
	return true
}

// Shuffle randomizes the order of cards in a Deck
func (d *Deck) Shuffle() {
	rand.Shuffle(d.Len(), func(i int, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

// Sort sorts the deck in place, using the specified ordering
func (d *Deck) Sort(order []Rank) {
	sort.Slice(d.Cards, func(i, j int) bool {
		iCard := d.Cards[i]
		jCard := d.Cards[j]
		iRankIndex := slices.Index(order, iCard.Rank)
		jRankIndex := slices.Index(order, jCard.Rank)
		return iRankIndex < jRankIndex
	})
}
