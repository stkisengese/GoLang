package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}
	prime := nb
	for !IsPrime(prime) {
		prime++
	}
	return prime
}
