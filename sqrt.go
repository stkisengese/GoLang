package piscine

func Sqrt(nb int) int {
	if nb == 0 {
		return 0
	}
	if nb == 1 {
		return 1
	}
	SqrtN := nb / 2
	for t := 1; t < 10; t++ {
		if SqrtN != 0 {
			SqrtN = SqrtN + (nb/SqrtN)/2
		} else {
			break
		}
		if SqrtN*SqrtN != nb {
			return 0
		}
	}
	return SqrtN
}
