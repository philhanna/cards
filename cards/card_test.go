package cards

import (
	"strings"
	"testing"
)

func TestCard_GetSVG(t *testing.T) {
	type fields struct {
		Rank Rank
		Suit Suit
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
	}{
		{"three of clubs", fields{THREE, CLUBS}, `sodipodi:docname="clubs_3.svg"`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Card{
				Rank: tt.fields.Rank,
				Suit: tt.fields.Suit,
			}
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
