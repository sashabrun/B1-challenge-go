package piscine

func ToUpper(s string) string {
	w := []rune(s)
	for i := 0; i < len(s); i++ {
		if s[i] >= 97 && s[i] <= 122 {
			w[i] = w[i] - 32
		}
	}
	return string(w)
}
