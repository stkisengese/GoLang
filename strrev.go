package piscine

import "github.com/01-edu/z01"

func StrRev(s string) string {
	Astr := []rune(s)
	for r, _ := range Astr {
		defer z01.PrintRune(Astr[r])
	}
	return string('\n')
}
