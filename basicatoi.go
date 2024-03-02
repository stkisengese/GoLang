package piscine

func BasicAtoi(s string) int {
	result := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			digit := int(c - '0')
			result = result*10 + digit
		} else {
			return 0
		}
	}
	return result
}
