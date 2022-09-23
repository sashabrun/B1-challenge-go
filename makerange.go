package piscine

func MakeRange(min, max int) []int {
	var liste []int
	if min > max {
		return liste
	}
	for i := min; i < max; i++ {
		liste = append(liste, i)
	}
	return liste
}
