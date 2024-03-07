package piscine

func BasicJoin(elems []string) string {
	var AstrJoin string
	for _, r := range elems {
		AstrJoin += r
	}
	return AstrJoin
}