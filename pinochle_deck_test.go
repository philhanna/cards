package cards

import (
	"sort"
	"testing"
)

func TestPinochleDeckSort(t *testing.T) {
	deck := make(PinochleDeck, 0)
	cards := []Card{
		{TEN, CLUBS},
		{JACK, DIAMONDS},
		{TEN, SPADES},
	}
	for _, card := range cards {
		deck = append(deck, card)
	}
	sort.Sort(&deck)
	card0 := Card{JACK, DIAMONDS}
	card1 := Card{TEN, CLUBS}
	card2 := Card{TEN, SPADES}
	if len(deck) != 3 {
		t.Errorf("len(deck): want=%d,have=%d", 3, len(deck))
	}
	if deck[0] != card0 {
		t.Errorf("card0: want=%v,have=%v", card0, deck[0])
	}
	if deck[1] != card1 {
		t.Errorf("card1: want=%v,have=%v", card1, deck[1])
	}
	if deck[2] != card2 {
		t.Errorf("card2: want=%v,have=%v", card2, deck[2])
	}

	
}
func TestPinochleDeck(t *testing.T) {
	deck := NewPinochleDeck()
	printDeck("New pinochle deck", []Card(*deck))
}

func TestPinochleDeckShuffled(t *testing.T) {
	deck := NewPinochleDeck()
	deck.Shuffle()
	printDeck("New pinochle deck shuffled", []Card(*deck))
}

func TestPinochleDeckShuffledThenSorted(t *testing.T) {
	deck := NewPinochleDeck()
	deck.Shuffle()
	sort.Sort(deck)
	printDeck("New pinochle deck shuffled then sorted", []Card(*deck))
}
