package piscine

func RecursivePower(nb int, power int) int {
	result := 1
	if power > 0 {
		result = nb * RecursivePower(nb, power-1)
	} else if power < 0 {
		return 0
	}
	return result
}
