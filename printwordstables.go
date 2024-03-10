package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, word := range a {
		for _, element := range word { // iterate through element runes
			z01.PrintRune(element)
		}
		z01.PrintRune('\n') // print new line after each word
	}
}
