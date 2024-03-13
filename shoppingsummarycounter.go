package piscine

import "strings"

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	shopping := strings.Fields(str)

	for _, item := range shopping {
		if count, ok := summary[item]; ok {
			summary[item] = count + 1
		} else {
			summary[item] = 1
		}
	}
	return summary
}
