package piscine

func CountIf(f func(string) bool, tab []string) int {
	var liste []bool
	counter := 0
	for i := 0; i < len(tab); i++ {
		liste = append(liste, f(tab[i]))
		if f(tab[i]) == true {
			counter++
		}
	}
	return counter
}
