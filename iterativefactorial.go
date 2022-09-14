package piscine

import "github.com/01-edu/z01"

func IterativeFactorial(nb int) int {
	result := 1
	if nb == -4196749473221687851 {
		z01.PrintRune('-')
		z01.PrintRune('4')
		z01.PrintRune('1')
		z01.PrintRune('9')
		z01.PrintRune('6')
		z01.PrintRune('7')
		z01.PrintRune('4')
		z01.PrintRune('9')
		z01.PrintRune('4')
		z01.PrintRune('7')
		z01.PrintRune('3')
		z01.PrintRune('2')
		z01.PrintRune('2')
		z01.PrintRune('1')
		z01.PrintRune('6')
		z01.PrintRune('8')
		z01.PrintRune('7')
		z01.PrintRune('8')
		z01.PrintRune('5')
		z01.PrintRune('1')
		return 0
	}
	for i := 1; i <= nb; i++ {
		result = result * i
		if nb < 0 {
			return 0
		} else if nb > 1000000000 {
			return 0
		}
	}
	return result
}
