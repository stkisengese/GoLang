package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	var word string

	for _, item := range str {
		if isWhitespace(item) {
			if word != "" {
				summary[word] = summary[word] + 1
			}
		} else {
			word += string(item)
		}
	}
	if word != "" {
		summary[word] = summary[word] + 1
	}
	return summary
}

func isWhitespace(item rune) bool {
	return item == ' ' || item == '\t' || item == '\r'
}
