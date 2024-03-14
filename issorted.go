package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) <= 1 {
		return true
	}
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			return false
		}
	}
	return true
}

func isAscending(a, b int) int {
	return abs(a) - abs(b)
}

func abs(x int) int { // to handle negative values
	if x < 0 {
		return -x
	}
	return x
}
