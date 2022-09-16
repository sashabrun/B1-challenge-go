package piscine

func Index(s string, toFind string) int {
	ind := 0
	indf := 0
	trouv := false
	if toFind == "" {
		return 0
	}
	if len(toFind) == 1 {
		for i := 0; i < len(s); i++ {
			if s[i] == toFind[0] {
				return i
			}
		}
	} else {
		for i := 0; i < len(s); i++ {

			if s[i] == toFind[0] {
				if !trouv {
					indf = i
				}
			}
			if s[i] == toFind[ind] {
				ind++
			} else {
				trouv = false
			}
			if ind > 1 && s[i] == toFind[len(toFind)-1] {
				return indf
			}
		}
	}
	return -1
}
