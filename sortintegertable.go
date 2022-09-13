package piscine

func SortIntegerTable(table []int) {
	n := len(table)
	b := false
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
				b = true
			}
		}
		if !b {
			return
		}
	}
}
