package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	good := true
	if n < 1 || n > 9 {
		return
	}
	num := 1
	for i := n; i > 0; i-- {
		num *= 10
	}
	for j := 0; j < num; j++ {
		good = true
		numberC := j
		for numberC >= 10 {
			if (numberC % 10) <= ((numberC / 10) % 10) {
				good = false
			}
			numberC /= 10
		}
		if good == true && j != 0 {
			fillZero(j, n)
			cpt := 1
			nombreCourrant := j
			for nombreCourrant >= 10 {
				cpt++
				nombreCourrant /= 10
			}
			if (numberC == 10-n) && (j%10) == 9 && cpt == n {
				z01.PrintRune('\n')
				return
			}
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
}

func PrintNbrbis(n int) {
	if n > 10 {
		PrintNbrbis(n / 10)
	}
	z01.PrintRune(rune((n % 10) + 48))
}
func fillZero(nb int, len int) {
	cpt := 1
	nombreCourrant := nb
	for nombreCourrant >= 10 {
		cpt++
		nombreCourrant /= 10
	}
	for cpt < len {
		z01.PrintRune(48)
		cpt++
	}
	PrintNbrbis(nb)
}
