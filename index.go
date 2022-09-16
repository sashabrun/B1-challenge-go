package piscine

func Index(s string, toFind string) int {
	trouve := true
	if len(toFind) == 0 {
		return 0
	}
	for i := 0; i < len(s); i++ {
		trouve = true
		if s[i] == toFind[0] {
			for x := 0; x < len(toFind); x++ {
				if s[i+x] != toFind[x] {
					trouve = false
				}
				if trouve == true {
					return i
				}
			}
		}
	}
	return -1
}
