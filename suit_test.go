package cards

import "testing"
import "github.com/stretchr/testify/assert"

func TestSuit_iterate(t *testing.T) {
	for i, suit := range Suits {
		switch i {
		case 0:
			assert.Equal(t, suit, SPADES)
		case 1:
			assert.Equal(t, suit, HEARTS)
		case 2:
			assert.Equal(t, suit, DIAMONDS)
		case 3:
			assert.Equal(t, suit, CLUBS)
		default:
			t.Errorf("%d is not one of the defined suits", i)
		}
	}
}
