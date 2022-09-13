package piscine

func Atoi(s string) int {
	j := 0
	if len(s) == 1 || len(s) == 0 {
		return 0
	}
	for i := 0; i < len(s); i++ {
		if 48 <= int(s[i]) && int(s[i]) <= 57 {
			j = j*10 + (int(s[i]) - 48)
		} else if (int(s[0]) == '+' && int(s[1]) == '+') || (int(s[0]) == '-' && int(s[1]) == '-') {
			return 0
		} else if int(s[i]) == 43 {
		} else if int(s[i]) == 45 {
		} else {
			return 0
		}
	}
	if int(s[0]) == 45 || int(s[1]) == 45 {
		j = -j
	}
	return j
}
