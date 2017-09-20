package goker

import (
	"time"
	"fmt"
)

// Creates all possible hands and evaluates them. Outputs the number of hands
// evaluated, the time elapsed, and the number of evaluations per second to
// the console
func Profile() {
	var hand Hand

	cnt := 0
	c1 := 0
	c2 := 1
	c3 := 2
	c4 := 3
	c5 := 4
	c6 := 5
	c7 := 6

	deck := MakeDeck()

	start := time.Now().UTC().UnixNano()

	for ; c1 < c2; c1++ {
		for c2 = c1+1; c2 < c3; c2++ {
			for c3 = c2+1; c3 < c4; c3++ {
				for c4 = c3+1; c4 < c5; c4++ {
					for c5 = c4+1; c5 < c6; c5++ {
						for c6 = c5+1; c6 < c7; c6++ {
							for c7 = c6+1; c7 < 52; c7++ {
								hand = Hand{}
								hand.AddCard(deck.GetByIndex(c7))
								hand.AddCard(deck.GetByIndex(c6))
								hand.AddCard(deck.GetByIndex(c5))
								hand.AddCard(deck.GetByIndex(c4))
								hand.AddCard(deck.GetByIndex(c3))
								hand.AddCard(deck.GetByIndex(c2))
								hand.AddCard(deck.GetByIndex(c1))
								hand.Evaluate()
								cnt++
							}
						}
					}
				}
			}
		}
	}

	end := time.Now().UTC().UnixNano()
	elapsed := float64(end - start) / float64(time.Second)

	fmt.Printf("Number of hands evaluated: %d\n", cnt)
	fmt.Printf("Time elapsed: %f seconds\n", elapsed)
	fmt.Printf("Evaluations per second: %d\n", int64(float64(cnt) / elapsed))
}
