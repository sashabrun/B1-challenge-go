package piscine

func Map(f func(int) bool, a []int) []bool {
	var liste []bool
	for i := 0; i < len(a); i++ {
		liste = append(liste, f(a[i]))
	}
	return liste
}
