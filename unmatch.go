package piscine

func Unmatch(a []int) int {
	for _, result := range a {
		count := 0
		for _, element := range a {
			if result == element {
				count++
			}
		}
		if count == 1 || count%2 == 1 {
			return result
		}
	}
	return -1
}
