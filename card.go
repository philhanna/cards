package cards

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// A Card is a combination of a Rank and a Suit
type Card struct {
	Rank Rank
	Suit Suit
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// Creates a new Card from a rank and suit
func NewCard(rank Rank, suit Suit) Card {
	p := new(Card)
	p.Rank = rank
	p.Suit = suit
	return *p
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Returns a representation of the card as a string
func (c Card) String() string {
	s := c.Rank.String() + c.Suit.String()
	return s
}

// Returns a representation of the card as a Unicode string
func (c Card) Unicode() string {
	suitOffset := c.Suit.Offset()
	rankOffset := c.Rank.Offset()	
	cardOffset := suitOffset + rankOffset
	s := fmt.Sprintf("%c", cardOffset)
	return s
}

// Returns the SVG image of this card
func (c Card) GetSVG() (string, error) {
	var rankName, suitName string

	// Get the rank name
	switch c.Rank {
	case TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN:
		rankName = fmt.Sprintf("%d", c.Rank)
	case JACK:
		rankName = "jack"
	case QUEEN:
		rankName = "queen"
	case KING:
		rankName = "king"
	case ACE:
		rankName = "ace"
	}

	// Get the suit name
	switch c.Suit {
	case SPADES:
		suitName = "spades"
	case HEARTS:
		suitName = "hearts"
	case DIAMONDS:
		suitName = "diamonds"
	case CLUBS:
		suitName = "clubs"
	}

	// Create the file name
	baseName := suitName + "_" + rankName + ".svg"
	fileName, _ := filepath.Abs(filepath.Join("svg_playing_cards", "fronts", baseName))

	// Load the SVG contents
	fp, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer fp.Close()
	contents, err := io.ReadAll(fp)
	if err != nil {
		return "", err
	}

	// Return contents as a string
	return string(contents), nil
}
