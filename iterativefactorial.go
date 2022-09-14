package piscine

func IterativeFactorial(nb int) int {
	result := 1
	for i := 1; i <= nb; i++ {
		result = result * i
	}
	if nb < 0 || nb > 100000000 {
		return 0
	} else {
		return result
	}
}
