package piscine

func IsPrintable(s string) bool {
	for _, r := range s {
		if r >= 32 && r <= 126 {
			continue
		} else {
			return false
		}
	}
	return true
}
