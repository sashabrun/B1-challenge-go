package piscine

func MakeRange(min, max int) []int {
	var liste []int
	if min >= max {
		return liste
	}
	liste2 := make([]int, max-min)
	valeur := min
	for i := 0; i < max-min; i++ {
		liste2[i] = valeur
		valeur++
	}
	return liste2
}
