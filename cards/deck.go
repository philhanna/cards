package cards

import "math/rand"

// -----------------------------------------------------------------------
// Type definitions
// -----------------------------------------------------------------------

// A Deck is an ordered list of cards.
type Deck struct {
	Cards   []Card
	Orderer RankOrder
}

// -----------------------------------------------------------------------
// Constructors
// -----------------------------------------------------------------------

// NewDeck constructs a 52-card regular deck and returns a pointer to it.
func NewDeck() *Deck {
	d := new(Deck)
	cards := make([]Card, 0)
	for i := 0; i < 4; i++ {
		suit := Suit(i)
		for j := 14; j >= 2; j-- {
			rank := Rank(j)
			card := NewCard(rank, suit)
			cards = append(cards, card)
		}
	}
	d.Cards = cards
	return d
}

// NewPinochleDeck constructs a 48-card Pinochle deck and returns a pointer to it.
func NewPinochleDeck() *Deck {
	d := new(Deck)
	cards := make([]Card, 0)
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
	d.Cards = cards
	return d
}

// -----------------------------------------------------------------------
// Methods
// -----------------------------------------------------------------------

// Shuffle randomizes the order of cards in a Deck
func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i int, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}
