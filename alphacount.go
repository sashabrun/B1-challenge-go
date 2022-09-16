package piscine

func AlphaCount(s string) int {
	var counter int
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			counter++
		}
	}
	return counter
}
