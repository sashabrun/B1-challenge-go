package piscine

func BasicAtoi2(s string) int {
	j := 0
	for i := 0; i < len(s); i++ {
		j = j*10 + (int(s[i]) - 48)
		if s[i] < 48 || s[i] > 57 {
			if s[i] >= 32 || s[i] <= 47 || s[i] >= 58 || s[i] <= 127 {
				return 0
			}
		}
	}
	return j
}
