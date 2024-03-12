package piscine

import "github.com/01-edu/z01"

func ForEach(f func(int), a []int) {
	for _, element := range a {
		f(element)
	}
}

func PrintNbr(n int) {
	digitRune := rune(n + '0')
	z01.PrintRune(digitRune)
}
