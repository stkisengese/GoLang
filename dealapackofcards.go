package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	players := 4
	cards := 12
	for i := 0; i < players; i++ {
		start := i * cards
		end := (i + 1) * cards

		playerCards := deck[start:end]

		fmt.Printf("Player %d: %v\n", players, playerCards)
	}
}
