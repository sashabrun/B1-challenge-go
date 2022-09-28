package piscine

func ForEach(f func(int), a []int) {
	for element := range a {
		f(element + 1)
	}
}
