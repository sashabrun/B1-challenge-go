package piscine

func Any(f func(string) bool, a []string) bool {
	var liste []bool
	for i := 0; i < len(a); i++ {
		liste = append(liste, f(a[i]))
		if f(a[i]) == true {
			return true
		}
	}
	return false
}
