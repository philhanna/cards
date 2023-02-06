package cards

import "testing"

func TestSuit_String(t *testing.T) {
	tests := []struct {
		name string
		suit Suit
		want string
	}{
		{"spades", SPADES, "\u2660"},
		{"hearts", HEARTS, "\u2661"},
		{"spades", DIAMONDS, "\u2662"},
		{"clubs", CLUBS, "\u2663"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.suit.String(); got != tt.want {
				t.Errorf("Suit.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
