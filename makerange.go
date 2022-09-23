package piscine

func MakeRange(min, max int) []int {
	var liste []int
	if min >= max {
		return liste
	}
	liste2 := make([]int, max-min)
	valeur := min
	for i := min; i < max; i++ {
		liste2[i] = valeur
		valeur++
	}
	return liste2
}
