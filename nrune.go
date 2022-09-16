package piscine

func NRune(s string, n int) rune {
	n -= 1
	i := []rune(s)
	if n < 0 || n >= len(s) {
		return 0
	}
	return i[n]
}
