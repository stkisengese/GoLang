package piscine

func IsPrintable(s string) bool {
	for _, r := range s {
		if r >= 'a' || r <= 'z' || r >= 'A' || r <= 'Z' {
			continue
		} else {
			return true
		}
	}
	return false
}
