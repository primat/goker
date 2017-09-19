package goker

import (
	"strconv"
	"math/bits"
)

type Card uint32
//type Card = uint32

//struct card{
//	unsigned num_A:2;
//	unsigned num_K:2;
//	unsigned num_Q:2;
//	//....
//	unsigned num_2:2;
//	unsigned spare:2;
//	unsigned spade:1;
//	unsigned heart:1;
//	unsigned diamond:1;
//	unsigned club:1;
//};

// Creates a single card from an integer.
// Useful for running in a loop to create a deck.
// i.e. 0 ==> 2 of Clubs, ...,  51 ==> Ace of Spades
func MakeCard(i int) Card {
	// Handle bad i values gracefully
	var ui uint32
	if i < 0 {
		ui = 0
	} else if i >= DECK_SIZE {
		ui = uint32(DECK_SIZE - 1)
	} else {
		ui = uint32(i)
	}
	return 1 << (2*(ui%13)+6) | 1 << (ui/13)
}

// Gets a simple text representation of a playing card
func (card Card) Text() string {

	b := ""
	val := bits.TrailingZeros32(uint32(card)>>6) / 2 + 2

	if val == 10 {
		b += "T"
	} else if val == 11 {
		b += "J"
	} else if val == 12 {
		b += "Q"
	} else if val == 13 {
		b += "K"
	} else if val == 14 {
		b += "A"
	} else {
		b += strconv.Itoa(val)
	}

	if card&1 == 1 {
		b += "C"
	} else if card&2 == 2 {
		b += "D"
	} else if card&4 == 4 {
		b += "H"
	} else if card&8 == 8 {
		b += "S"
	}
	//b += "♣"
	//b += "♦"
	//b += "♥"
	//b += "♠"

	return b
}
