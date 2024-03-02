package piscine

func StrRev(s string) string {
	Astr := []rune(s)
	var revRun []rune
	for i := len(Astr) - 1; i >= 0; i-- {
		revRun = append(revRun, (Astr[i])) 	
	}	
}
