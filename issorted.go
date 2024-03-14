package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	sorted := true
	length := len(a)
	for i := 0; i < length-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			sorted = false
		}
	}
	if !sorted {
		sorted = true
		for i := 0; i < length-1; i++ {
			if f(a[i], a[i+1]) < 0 {
				sorted = false
			}
		}
	}
	return sorted
}
