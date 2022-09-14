package piscine

func RecursiveFactorial(nb int) int {
	result := 1
	if nb < 0 {
		return 0
	} else if nb > 1000000000 {
		return 0
	}
	if i := 1; i <= nb {
		i++
		result = result * i
	}
	return result
}
