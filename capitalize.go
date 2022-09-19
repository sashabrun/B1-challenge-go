package piscine

func Capitalize(s string) string {
	res := ""
	maj := false
	for i := 0; i < len(s); i++ {
		if i == 0 {
			maj = true
		}

		if maj {
			if s[i] >= 97 && s[i] <= 122 {
				res += string(s[i] - 32)
			} else {
				res += string(s[i])
			}
			maj = false
		} else {
			if s[i] >= 'A' && s[i] <= 'Z' {
				res += string(s[i] + 32)
			} else {
				res += string(s[i])
			}
		}

		if (s[i] < 'A' || s[i] > 'Z') && (s[i] < 'a' || s[i] > 'z') && (s[i] < '0' || s[i] > '9') {
			maj = true
		}
	}
	return res
}
