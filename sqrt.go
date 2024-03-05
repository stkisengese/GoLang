package piscine

func Sqrt(nb int) int {
	for i := 0; 1 <= nb; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}
