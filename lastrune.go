package piscine

func LastRune(s string) rune {
  Astr := [] rune(s)  // slice the runes
  if len(Astr) == 0 {
    return 0
  } 
  return Astr[len(Astr)-1]
}
