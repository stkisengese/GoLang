package piscine

func Capitalize(s string) string {
	var capital string
	measure := true
	for _, r := range s {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
			if measure {
				if r >= 'a' && r <= 'z' {
					capital += string(r - 32)
				} else {
					capital += string(r)
				}
				measure = false
			} else {
				if r >= 'A' && r <= 'Z' {
					capital += string(r + 32)
				} else {
					capital += string(r)
				}
			}
		} else {
			capital += string(r)
			measure = true
		}
	}
	return capital
}