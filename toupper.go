package piscine

func ToUpper(s string) string {
	Astr := ""
	for _, r := range s {
		if r >= 97 && r <= 122 {
			r = (r - 32)
		}
		Astr = Astr + string(r)
	}
	return Astr
}
