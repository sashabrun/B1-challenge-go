package piscine

func Swap(a *int, b *int) {
	r := *b
	*b = *a
	*a = r
}
