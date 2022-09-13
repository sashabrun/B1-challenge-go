package piscine

func BasicAtoi(s string) int {
	j := 0
	for i := 0; i < len(s); i++ {
		j = j*10 + (int(s[i]) - 48)
	}
	return j
}
