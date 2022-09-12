package piscine

func StrRev(s string) string {
	resultat := ""
	for i := 0; i < len(s); i++ {
		resultat = string(s[i]) + resultat
	}
	return resultat
}
