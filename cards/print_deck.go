package cards

import (
	"fmt"
	"strings"
)

func printDeck(label string, cards []Card) {
	fmt.Printf("%s:\n", label)
	buffer := make([]string, 0)
	n := len(cards) / 4
	for _, card := range cards {
		buffer = append(buffer, card.Unicode())
		if len(buffer) == n {
			fmt.Printf("%s\n", strings.Join(buffer, " "))
			buffer = buffer[0:0]
		}
	}
}
