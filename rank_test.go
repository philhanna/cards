package cards

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRank_Compare(t *testing.T) {
	const (
		LESS    = -1
		GREATER = 1
		EQUAL   = 0
	)
	tests := []struct {
		name  string
		r     Rank
		other Rank
		want  int
	}{
		{"3 > 2", THREE, TWO, GREATER},
		{"J = J", JACK, JACK, EQUAL},
		{"10 < K", TEN, KING, LESS},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Compare(tt.other); got != tt.want {
				t.Errorf("Rank.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRank_ComparePinochle(t *testing.T) {
	const (
		LESS    = -1
		GREATER = 1
		EQUAL   = 0
	)
	tests := []struct {
		name  string
		r     Rank
		other Rank
		want  int
	}{
		{"9 < Q", NINE, QUEEN, LESS},
		{"J = J", JACK, JACK, EQUAL},
		{"10 > K", TEN, KING, GREATER},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.ComparePinochle(tt.other); got != tt.want {
				t.Errorf("Rank.ComparePinochle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRank_String(t *testing.T) {
	tests := []struct {
		name string
		r    Rank
		want string
	}{
		{"numbered card", FIVE, "5"},
		{"Jack", JACK, "J"},
		{"Queen", QUEEN, "Q"},
		{"King", KING, "K"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			have := tt.r.String()
			assert.Equal(t, want, have)
		})
	}
}
