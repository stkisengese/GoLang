package piscine

func JumpOver(str string) string {
	var finalStr string
	if len(str) > 3 {
		for index, value := range str {
			if index%3 == 2 {
				finalStr += string(value)
			}
		}
	}
	finalStr += "\n"
	return finalStr
}
