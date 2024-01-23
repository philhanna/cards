package cards

import "math/rand"
import "slices"

// -----------------------------------------------------------------------
// Type definitions
// -----------------------------------------------------------------------

// Deck is the abstract base type for RegularDeck and PinochleDeck
type Deck struct {
	Cards []Card
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewDeck returns a new Deck, possibly with an initial set of cards
func NewDeck(card ...Card) Deck {
	return Deck{
		Cards: card,
	}
}

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// Compare compares the ranks of two cards and returns a negative
// integer if the first card has a lower rank, a positive integer if the
// first card has a higher rank, or zero, if both have the same rank.
func Compare(iCard Card, jCard Card, ranks []Rank) int {
	iRankIndex := slices.Index(ranks, iCard.Rank)
	jRankIndex := slices.Index(ranks, jCard.Rank)
	return iRankIndex - jRankIndex
}

// -----------------------------------------------------------------------
// Methods
// -----------------------------------------------------------------------

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
	rand.Shuffle(len(d.Cards), func(i int, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

// -----------------------------------------------------------------------
// Implementation of the sort interface. RegularDeck and PinochleDeck
// have their own Less() methods.
// -----------------------------------------------------------------------

// Len returns the number of cards in the deck.
func (d Deck) Len() int {
	return len(d.Cards)
}

// Exchanges two Rank objects in the array.
func (d Deck) Swap(i int, j int) {
	d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
}
