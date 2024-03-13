package piscine

func StringToIntSlice(str string) []int {
	var intSlice []int
	for _, char := range str {
		intSlice = append(intSlice, int(char))
	}
	return intSlice
}
