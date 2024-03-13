package piscine

import "fmt"

func DealPackOfCards(deck []int) {
	if len(deck) != 12 {
		return
	}
	cardsPerPlayer := len(deck) / 4
	for player := 1; player <= 4; player++ {
		start := (player - 1) * cardsPerPlayer
		end := start + cardsPerPlayer
		playerCards := deck[start:end]

		fmt.Printf("Player %d: %v\n", player, playerCards)
	}
}
