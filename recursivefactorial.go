package piscine

func RecursiveFactorial(nb int) int {
	if nb == 1 || nb == 0 {
		return nb
	}
	return nb * (RecursiveFactorial(nb - 1))
}
