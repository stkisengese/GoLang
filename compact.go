package piscine

func Compact(ptr *[]string) int {
	res := *ptr
	count := 0
	for _, value := range *ptr {
		if value != "" {
			res[count] = value
			count++
		}
	}
	*ptr = res[0:count]
	return count
}
