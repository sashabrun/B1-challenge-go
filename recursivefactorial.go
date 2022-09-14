package piscine

func RecursiveFactorial(nb int) int {
	if nb == 1 {
		return nb
	} else if nb > 50 || nb == 0 {
		return 0
	}
	return nb * (RecursiveFactorial(nb - 1))
}
