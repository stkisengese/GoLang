package piscine

func Abort(a, b, c, d, e int) int {
	number := []int{a, b, c, d, e}
	for i := 0; i < len(number)-1; i++ {
		for j := 0; j < len(number)-1; j++ {
			if number[j] > number[j+1] {
				f := number[j]
				number[j] = number[j+1]
				number[j+1] = f
			}
		}
	}
	return number[2]
}
