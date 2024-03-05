package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for {
		if IsPrime(nb) {
			return true
		} else {
			return false
		}
	}
}
