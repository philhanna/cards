package cards

import (
	"sort"
	"testing"
)

func TestDeckSort(t *testing.T) {
	deck := make(Deck, 0)
	cards := []Card{
		{TEN, CLUBS},
		{JACK, DIAMONDS},
		{TEN, SPADES},
	}
	for _, card := range cards {
		deck = append(deck, card)
	}
	sort.Sort(&deck)
	card0 := Card{TEN, CLUBS}
	card1 := Card{TEN, SPADES}
	card2 := Card{JACK, DIAMONDS}
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

func TestDeck(t *testing.T) {
	deck := NewDeck()
	printDeck("New deck", []Card(*deck))
}

func TestDeckShuffled(t *testing.T) {
	deck := NewDeck()
	deck.Shuffle()
	printDeck("Shuffled deck", []Card(*deck))
}

func TestDeckShuffledThenSorted(t *testing.T) {
	deck := NewDeck()
	deck.Shuffle()
	sort.Sort(deck)
	printDeck("Shuffled then sorted deck", []Card(*deck))
}
