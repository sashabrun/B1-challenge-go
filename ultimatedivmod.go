package piscine

func UltimateDivMod(a *int, b *int) {
	r := *a / *b
	*b = *a % *b
	*a = r
}
