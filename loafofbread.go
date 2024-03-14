package piscine

func LoafOfBread(str string) string {
	result := ""
	i := 0
	for i < len(str) {
		if i+5 > len(str) {
			return "Invalid Output\n"
		}
		substring := ""
		for j := 0; j < 5; j++ {
			if str[i+j] == ' ' {
				continue
			}
			substring += string(str[i+j])
		}
		result += substring + "\n"
		i += 6
	}
	if len(result) > 0 && result[len(result)-1] == '\n' {
		result = result[:len(result)-1] // slice to remove the last character
	}
	return result
}
