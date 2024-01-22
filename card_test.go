package cards

import (
	"strings"
	"testing"
)

func TestCard_GetSVG(t *testing.T) {
	tests := []struct {
		name string
		card Card
		want string
	}{
		{"three of clubs", NewCard(THREE, CLUBS), `sodipodi:docname="clubs_3.svg"`},
		{"ace of spades", NewCard(ACE, SPADES), `sodipodi:docname="spades_ace_simple.svg"`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.card
			have, err := c.GetSVG()
			if err != nil {
				t.Error(err)
			}
			ok := strings.Contains(have, tt.want)
			if !ok {
				t.Errorf("Does not contain %s", tt.want)
			}
		})
	}
}
