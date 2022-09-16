package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n >= 0 {
		var liste []int
		if n == 0 {
			liste = append(liste, 0)
		}
		for n != 0 {
			if n < 0 {
				n = -n
			}
			liste = append(liste, n%10)
			n = n / 10
		}
		SortIntegerTable2(liste)
		for i := 0; i < len(liste); i++ {
			z01.PrintRune(rune(liste[i]) + 48)
		}
	}
}

func SortIntegerTable2(table []int) {
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
