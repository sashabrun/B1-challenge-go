package piscine

func TrimAtoi(s string) int {
	j := 0
	neg := false
	signe := true
	for i := 0; i < len(s); i++ {
		if 48 <= int(s[i]) && int(s[i]) <= 57 {
			j = j*10 + (int(s[i]) - 48)
			signe = false
		}
		if signe {
			if s[i] == 45 {
				neg = true
			}
		}
	}
	if neg {
		j = -j
	}
	return j
}
