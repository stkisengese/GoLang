package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}
	prime := nb
	for {
		if IsPrime(prime) {
			return prime
		}
		prime++
	}
}
