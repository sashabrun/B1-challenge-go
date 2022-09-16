package piscine

func ToLower(s string) string {
	w := []rune(s)
	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 {
			w[i] = w[i] + 32
		}
	}
	return string(w)
}
