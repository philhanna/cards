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
	cards []Card
}

// -----------------------------------------------------------------------
// Methods
// -----------------------------------------------------------------------

// Len returns the number of cards in the deck
func (d *Deck) Len() int {
	return len(d.cards)
}

// Remove removes the specified card from the deck, if it exists.
// Returns true if the card was removed.
func (d *Deck) Remove(card Card) bool {
	p := slices.Index(d.cards, card)
	if p == -1 {
		return false
	}
	oldCards := d.cards[:]
	d.cards = append(oldCards[:p], oldCards[p+1:]...)
	return true
}

// Shuffle randomizes the order of cards in a Deck
func (d *Deck) Shuffle() {
	rand.Shuffle(d.Len(), func(i int, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

// Sort sorts the deck in place, using the specified ordering
func (d *Deck) Sort(order []Rank) {
	sort.Slice(d.cards, func(i, j int) bool {
		iCard := d.cards[i]
		jCard := d.cards[j]
		iRankIndex := slices.Index(order, iCard.Rank)
		jRankIndex := slices.Index(order, jCard.Rank)
		return iRankIndex < jRankIndex
	})
}
