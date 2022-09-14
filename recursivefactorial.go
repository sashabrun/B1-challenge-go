package piscine

func RecursiveFactorial(nb int) int {
	if nb == 1 || nb == 0 {
		return nb
	} else if nb > 100000000 {
		return 0
	}
	return nb * (RecursiveFactorial(nb - 1))
}
