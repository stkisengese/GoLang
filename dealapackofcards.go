package piscine

import "github.com/01-edu/z01"

func DealAPackOfCards(deck []int) {
	players := 4
	cards := len(deck) / 4
	for i := 0; i < players; i++ {
		start := i * cards
		end := (i + 1) * cards
		if i == players-1 {
			end = len(deck)
		}
		playerCards := deck[start:end]
		z01.PrintRune("Player ")
		z01.PrintRune(i + 1)
		z01.PrintRune(':')
		z01.PrintRune(' ')
		z01.PrintRune(playerCards)
		z01.PrintRune('\n')
	}
}
