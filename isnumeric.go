package piscine

func IsNumeric(s string) bool {
	for _, r := range s {
		if r >= '0' && r <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
