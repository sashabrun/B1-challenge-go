package piscine

func Split(s, sep string) []string {
	var liste []string
	element := ""
	i := 0
	for i < len(s) {
		if s[i] == sep[0] {
			trouv := false
			x := i
			for j := 1; j < len(sep); j++ {
				if s[x+j] != sep[j] {
					element += string(s[i])
				} else if j == len(sep)-1 {
					trouv = true
					i++
				} else {
					i++
				}
			}
			if trouv {
				liste = append(liste, element)
				element = ""
			}
			i++
		} else {
			element = element + string(s[i])
			i++
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
