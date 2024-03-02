package piscine

func StrRev(s string) string {
	Astr := []rune(s)

	for i, j := 0, (len(Astr) - 1); i > j; i, j = i+1, j-1 {
		Astr[i], Astr[j] = Astr[j], Astr[i]
	}
	return string(Astr)
}
