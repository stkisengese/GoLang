package piscine

func NRune(s string, n int) rune {
	Astr := []rune(s)
	if n <= 0 || n > len(Astr) {
		return 0
	}
	return Astr[n-1]
}
