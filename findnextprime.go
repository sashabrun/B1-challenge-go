package piscine

func FindNextPrime(nb int) int {
	for i := nb; i < 432634643; i++ {
		if IsPrime(i) {
			return i
		}
	}
	return 0
}
