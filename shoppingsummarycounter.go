package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	receipt := make(map[string]int)
	var bought string

	for _, item := range str {
		if item == 32 {
			receipt[bought] += 1
			bought = ""
		} else if item != 32 {
			bought += string(byte(item))
		}
	}
	receipt[bought] += 1

	return receipt
}
