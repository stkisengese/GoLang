package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	if n < 0 {
		return
	}
	for digit := 0; digit < 10; digit++ {
		counter := 0
		for k := n; k > 0; k /= 10 {
			if k%10 == digit {
				counter++
			}
		}
		for i := 0; i < counter; i++ {
			z01.PrintRune(rune(digit + '0'))
		}
	}
}
