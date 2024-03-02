package piscine

func StrRev(s string) string {
	Astr := []rune(s)
	for i := 0; i < len(Astr)-1; i += 2 {
		Astr[i], Astr[i+1] = Astr[i+1], Astr[i]
	}
	return string(Astr)
}
