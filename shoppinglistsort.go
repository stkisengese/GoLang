package piscine

func ShoppingListSort(slice []string) []string {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if len(slice[i]) > len(slice[j]) {
				swap := slice[i]
				slice[i] = slice[j]
				slice[j] = swap
			}
		}
	}
	return slice
}
