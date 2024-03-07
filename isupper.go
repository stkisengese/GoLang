package piscine

func IsUpper(s string) bool {
	n := 0
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			n++
		}
	}
	return n == len(s)
}
