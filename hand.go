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
	// This data structure doesn't hold the Card types, rather it stores the
	// sum of the cards for each suit
	// cards[0] counts spades
	// cards[1] counts clubs
	// cards[2] counts diamonds
	// cards[3] should always contain the other counts OR'ed together
	// cards[4] counts hearts
}


// Compresses a 26-bit card representation to a 13-bit card representation
func (h *Hand) compress(a uint32) uint32 {
	var i uint32 = 0
	var out uint32 = 0

	for ; a > 0; i++ {
		out |= (a&1) << uint32(i)
		a /= 4
	}

	return out/8
}

// Add a card to the hand.
func (h *Hand) AddCard(card Card) {
	h.cards[card&7] += uint32(card)
	h.cards[3] |= uint32(card)
}

// Evaluates a 7, 6, or 5 card poker hand
func (h *Hand) Evaluate() uint32 {
	// Variables used:
	// count: the sum of all suits. counts the ranks in parallel.
	//  But there are only 2 bits used to store each rank, so 4 of a kind will
	//  overflow. We fix that by subtracting h.cards[3] which has a 1 for each
	//  rank actually in the hand. So now every 2-bit field holds the count-1
	//  for that rank.
    // evens: it has a bit set in any rank which has a 1(pair) or 3(quad).
    // odds: it has a bit for every 2(set) or 3(quad).
    // result: Holds the type of hand:
	// 	9 = Straight Flush
	//  7 = Quad (4 of a kind)
	//  6 = Boat (full house)
	//  5 = Flush
	//  4 = Run (straight)
	//  3 = Set (3 of a kind)
	//  2 = 2 pair
	//  1 = Pair
	//  0 = High card
	// value: The cards that determine the hand value
    // kicker: Holds the tiebreak card(s) (eg: KKA21 vs KKQ21)

	count := h.cards[0] + h.cards[1] + h.cards[2] + h.cards[4] - (h.cards[3] & 0xFFFFFFF0)
	evens := 0x55555540 & count
	odds := 0xAAAAAA80 & count

	var (
		value uint32
		result uint32 = 0
		kicker uint32 = h.cards[3]
	)

	// Quad detector:
	// The value `v=e&o/2` will be non-zero only if a rank has both the even
	// and odd bit set - meaning its count-1 is 3. The while loop clears all
	// except the top bit of the remaining to find the kicker
	if value = evens & (odds / 2); value != 0 {
		//fmt.Println("Four of a kind!")
		result = 7
		kicker = h.cards[3] ^ value
		temp := kicker & (kicker - 1)
		for temp != 0 {
			kicker = temp
			temp = kicker & (kicker - 1)
		}

	// Full House detector:
	// The first else block catches 2 sets (odds counter has 2 bits set).
	// It separates the bits into high set in `value` and the pair in `kicker`
	// The second else block catches a set and one or two pairs.
	// When setting `kicker`, it clears one bit from the pairs field if needed
	// (since AAAKKQQ ranks the same as AAAKKQJ)
	} else if value = odds & (odds-1); value != 0 {
		result = 6
		value /= 2;
		kicker = (odds/2) ^ value

	} else if evens != 0 && odds != 0 {
		result = 6
		value = odds / 2
		temp := evens & (evens-1)
		if temp != 0 {
			kicker = temp
		} else {
			kicker = evens
		}
	} else {

		// All other hands fall here.
		// h.cards[3] is in `kicker`. It will be used to detect straights. (it
		// holds a bit for each unique value and a bit for each unique suit)

		// Flush detector:
		// Remember that for suit X=1,2,4,8: h[X&7] holds a 3-bit card count,
		// Starting at bit 0,1,2,3 respectively, subtract 1 from the count and
		// store in `C` If C > 4, there is a five card flush.
		// Overwrite `kicker` with the flush suit, since a plain straight won't
		// beat this, but a straight flush will.
		for i := 0; i < 4; i++ {
			idx := (1 << uint32(i)) & 7
			count = h.cards[idx] >> uint32(i)
			count &= 7
			if count >= 5 {
				kicker = h.cards[idx]
				result = 5
				break
			}
		}

		// Straight detector:
		// Clear the suit bits from a, then copy down the high bit (ace)
		// to the ones position so we can catch 5-high straights.
		kicker &= 0xFFFFFFC0
		value = kicker | ((kicker >> 26) & 16)

		// The next line zeros value unless there are at least 5 cards in a
		// row. `result` will be 4 for straights, 9 for straight flushes.
		// For a 6 or 7 card straight, there will be multiple consecutive bits
		// set in value: `value &= ^(value/4)` clears all but the highest.
		for i := 0; i < 4; i++ {
			value &= value * 4
		}

		var i uint32

		if value != 0 {
			result += 4
			value &= ^(value / 4)
			kicker = value

		// Finish up the flush detection:
		// 'result' is only set for flush, store the high 5 cards in `kicker`
		// and `value`, by clearing the low bit until the card count `count`
		// is 5 (done after straight detection to avoid calling AK98765 in same
		// suit a plain flush.) `i` will be 0 for cases below here.
		} else if  i = result; i != 0 {
			for ; count > 5; count-- {
				kicker &= kicker-1 // k^v has 0 bits, i does not matter
			}
			value = kicker

		// Three of a kind detector:
		// Two sets are a full house, caught above. So if there is any bit left
		// in 'odds', it is a set. `v = o/2` shifts the value bit into the
		// right place
		} else if value = odds / 2; value != 0 {
			result = 3

		// Pairs detector:
		// A bit set in evens is a pair. There may exist 1, 2, or 3 of them.
		// `odds` will be set if there is more than one. `i` will be set if
		// there are 3. `value` is set to the top 1 or 2 cards. 't' is 1 or 2.
		} else if evens != 0 {
			odds = evens & (evens-1)
			i = odds & (odds-1)
			if i != 0 {
				value = odds
			} else {
				value = evens
			}
			result = 1
			if odds > 0 {
				result++
			}
		}

		// For all hands except 4 of a kind and full house,
		// we have left the primary cards which determine the hand's type in
		// 'value' and `a` holds all the cards (except a == v for flushes and
		// straights.) Set k to the kickers by finding all in `a` not in
		// `value` (a^v) then clear the extra 2. (Or 1 if i is non zero b/c
		// there was a 3rd pair).
		kicker ^= value
		kicker &= kicker-1
		if i==0 {
			kicker &= kicker-1
		}
	}

	// Build the final result.
	// 4 bits for the type 0..9, 13 bits for the value cards, 13 for the kicker
	value = h.compress(value);
	C := h.compress(kicker);
	return (result<<28) | (value<<13) | C
}

// Outputs cards in the deck to the console - used for development
func (h *Hand) ToConsole() {
	for i := range h.cards {
		fmt.Printf("%-2v %032v\n", i, strconv.FormatUint(uint64(h.cards[i]), 2) )
	}
}
