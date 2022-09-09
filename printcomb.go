package main

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 48; i < 58; i++ {
		for x := 48; x < 58; x++ {
			for k := 48; k < 58; k++ {
				if i < x && x < k {
					z01.PrintRune(rune(i))
					z01.PrintRune(rune(x))
					z01.PrintRune(rune(k))
					if i != 55 {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
