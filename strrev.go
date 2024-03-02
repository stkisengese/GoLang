package piscine

import "github.com/01-edu/z01"

func StrRev(s string) string {
	Astr := []rune(s)
	for i := len(Astr) - 1; i >= 0; i-- {
		z01.PrintRune(Astr[r])
	}
	return string('\n')
}