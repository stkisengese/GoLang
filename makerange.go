package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	result := make([]int, max-min)
	for i := min; i < max; i++ {
		result[i] = i
	}
	return result
}
