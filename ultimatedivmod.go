package piscine

func UltimateDivMod(a *int, b *int) {
	Div := *a / *b
	Mod := *a % *b
	*a = Div
	*b = Mod
}
