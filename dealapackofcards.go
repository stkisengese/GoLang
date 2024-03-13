package piscine

import "fmt"

func DealPackOfCards(deck []int) {
	players := 4
	cards := len(deck) / 4
	for i := 0; i < players; i++ {
		start := i * cards
		end := (i + 1) * cards

		playerCards := deck[start:end]

		fmt.Printf("Player %d: %v\n", players, playerCards)
	}
}
