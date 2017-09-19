package goker

import (
	"fmt"
	"strconv"
)

// The Hand is five 32-bit words. There's one word for each suit, and one
// which has the suit words OR'ed together, leaving a bit set for each card
// value present.
type Hand struct {
	cards [5]uint32
	// This data structure doesn't hold the Card types, rather it stores the sum of the cards for each suit
	// cards[0] counts spades
	// cards[1] counts clubs
	// cards[2] counts diamonds
	// cards[3] should always contain the other counts OR'ed together
	// cards[4] counts hearts

}

// Add a card to the hand.
func (h *Hand) AddCard(card Card) {
	h.cards[card&7] += uint32(card)
	h.cards[3] |= uint32(card)
}

//
// Outputs cards in the deck to the console - used in development
func (h *Hand) ToConsole() {
	for i := range h.cards {
		fmt.Printf("%-2v %032v\n", i, strconv.FormatUint(uint64(h.cards[i]), 2) )
	}
}
