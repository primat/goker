package main

import (
	"math/rand"
	"time"
	"github.com/primat/goker"
	"fmt"
)

//
// Application entry point
func main() {

	// Set a random seed unless we want the same results on each execution
	rand.Seed(time.Now().UTC().UnixNano())

	// Create a deck of cards
	//deck := goker.MakeDeck()
	deck := goker.MakeShuffledDeck()
	//deck.ToConsole()

	// Create an empty hand
	var hand = goker.Hand{}

	// Deal 7 cards to the hand
	for i := 0; i < 7; i++ {
		c := deck.Pop()
		fmt.Println(c.Text())
		hand.AddCard(c)
	}

	hand.ToConsole()
	
}

//func ACE_makecard(int i){return 1<<(2*(i%13)+6)|1<<(i/13);}

