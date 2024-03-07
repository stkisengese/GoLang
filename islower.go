package piscine

func IsLower(s string) bool {
	n := 0
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			n++
		}
	}
	return n == len(s)
}
