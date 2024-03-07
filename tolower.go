package piscine

func ToLower(s string) string {
	var Astr string
	for _, r := range s {
		if r >= 65 && r <= 90 {
			r = (r + 32)
		}
		Astr += string(r)
	}
	return Astr
}
