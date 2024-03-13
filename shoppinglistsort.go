package piscine

func ShoppingListSort(slice []string) []string {
	for i := 0; i < len(slice)-1; i++ { // loop through the slice
		for j := i + 1; j < len(slice); j++ { // loop through remaining elements
			if len(slice[i]) > len(slice[j]) { // compare string length
				swap := slice[i] // swap elements if the length of i is greater
				slice[i] = slice[j]
				slice[j] = swap
			}
		}
	}
	return slice
}
