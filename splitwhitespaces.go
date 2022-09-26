package piscine

func SplitWhiteSpaces(s string) []string {
	var liste []string
	element := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			liste = append(liste, element)
			element = ""
		}
		element = element + string(s[i])
	}
	liste = append(liste, element)
	return liste
}
