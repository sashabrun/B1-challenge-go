package piscine

func IterativeFactorial(nb int) int {
	result := 1
	for i := 1; i <= nb; i++ {
		result = result * i
		if nb < 0 {
			return 0
		} else if nb > 1000000000 {
			return 0
		}
	}
	return result
}
