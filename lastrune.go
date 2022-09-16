package piscine

func LastRune(s string) rune {
	i := []rune(s)
	return i[len(s)-1]
}
