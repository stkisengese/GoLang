package piscine

func FirstRune(s string) rune {
	for _, r := range s {
		return r
	}
	return 0
}
// str := []rune (s)
// return str[len(str)-1]
// return str[0]