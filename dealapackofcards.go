package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	cards := 4
	k := 0
	for i := 1; i <= cards; i++ {
		fmt.Printf("Player %v: %d, %d, %d\n", i, deck[k], deck[k+1], deck[k+2])
		k += 3
	}
}
