package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	/*if n < 0 {
		z01.PrintRune(rune('0' + -n%10))
		z01.PrintRune(rune('0' + (-n/10)%10))
		z01.PrintRune(rune('0' + (-n/100)%10))
	} else if n > 0 {
		z01.PrintRune(rune('0' + n%10))
		z01.PrintRune(rune('0' + (n/10)%10))
		z01.PrintRune(rune('0' + (n/100)%10))
	} else {
		z01.PrintRune(rune('0'))
	}*/
	if n > 10 {
		PrintNbr(n / 10)
	} else if n < 0 {
		z01.PrintRune(rune(45))
		n = -n
		PrintNbr(n / 10)
	}
	z01.PrintRune(rune(*/.n + 48))
}
