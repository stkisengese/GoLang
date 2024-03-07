package piscine

func TrimAtoi(s string) int {
	number := 0
	symbol := 1
	for _, r := range s {
		if r >= '0' && r <= '9' {
			number = number*10 + int(r-'0')
		} else if r == '-' && number == 0 {
			symbol = -1
		} else if number > 0 {
			continue
		}
	}
	return symbol * number
}
