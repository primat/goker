package goker

import (
	"math/rand"
	"fmt"
	"strconv"
)

const DECK_SIZE int = 52

// A deck is basically an empty array with capacity 'DECK_SIZE' when
// initialized
type deck struct {
	Cards []Card
}


// Creates a deck with all cards ordered
func MakeDeck() *deck {
	deck := deck{}
	deck.init()

	for i := 0; i < DECK_SIZE; i++ {
		deck.Cards = append(deck.Cards, MakeCard(i))
	}

	return &deck
}

// Creates a deck of randomly ordered cards
func MakeShuffledDeck() *deck {
	deck := deck{}
	deck.init()
	randomList := rand.Perm(DECK_SIZE)

	for i := range randomList {
		deck.Cards = append(deck.Cards, MakeCard(randomList[i]))
	}

	return &deck
}

// Common operations for initializing a new deck
func (d *deck) init() {
	d.Cards = make([]Card, 0, DECK_SIZE)
}

// Gets a card by index (used for testing with a sorted deck)
func (d *deck) GetByIndex(i int) Card {
	return d.Cards[i]
}

// Pops a card off the top of the deck
func (d *deck) Pop() Card {
	var c Card
	c, d.Cards = d.Cards[len(d.Cards)-1], d.Cards[:len(d.Cards)-1]

	return c
}

// Returns the number of cards currently in the deck
func (d *deck) Size() int {
	return len(d.Cards)
}

// Outputs cards in the deck to the console - used in development
func (d *deck) ToConsole() {
	for i := range d.Cards {
		fmt.Printf("%-2v %032v %s\n", i+1, strconv.FormatUint(uint64(d.Cards[i]), 2), d.Cards[i].Text() )
	}
}
