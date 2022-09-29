package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	c := 0
	d := 0
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) >= 0 {
			d++
		}
		if f(a[i-1], a[i]) <= 0 {
			c++
		}
	}
	if d == len(a)-1 {
		return true
	} else if c == len(a)-1 {
		return true
	} else {
		return false
	}
}
