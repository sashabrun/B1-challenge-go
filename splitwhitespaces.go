package piscine

func SplitWhiteSpaces(s string) []string {
	var liste []string
	element := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			liste = append(liste, element)
			element = ""
		} else {
			element = element + string(s[i])
		}
	}
	liste = append(liste, element)
	var liste2 []string
	for i := 0; i < len(liste); i++ {
		if liste[i] != "" {
			liste2 = append(liste2, liste[i])
		}
	}
	return liste2
}
