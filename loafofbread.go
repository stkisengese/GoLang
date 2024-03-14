package piscine

func LoafOfBread(str string) string {
	result := ""
	for i := 0; i < len(str); i++ {
		if i+5 > len(str) {
			return ""
		}
		word := ""
		for j := 0; j < 5 && i+j < len(str); j++ {
			if str[i+j] != ' ' {
				word += string(str[i+j])
			}
		}
		if word != "" || i+len(word) == len(str) {
			result += word + "\n"
		}
		i += len(word)

	}
	if len(result) > 0 && result[len(result)-1] == '\n' {
		result = result[:len(result)-1] // slice to remove the last character
	}
	return result
}
