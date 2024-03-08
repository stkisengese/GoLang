package piscine

func Join(strs []string, sep string) string {
	var AstrAdd string
	for i, r := range strs {
		AstrAdd += r
		if i < len(strs)-1 {
			AstrAdd += sep
		}
	}
	return AstrAdd
}
