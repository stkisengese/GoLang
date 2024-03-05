package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}
	prime := 1
	for i := 2; i < nb; i++ {
		if (nb % i) == 0 {
			nb++
		}
	}
	if prime == 1 {
		return nb
	} else {
		return nb
	}
}
