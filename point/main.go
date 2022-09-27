package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x rune
	y rune
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	z01.PrintRune((points.x / 10) + 48)
	z01.PrintRune((points.x % 10) + 48)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	z01.PrintRune((points.y / 10) + 48)
	z01.PrintRune((points.y % 10) + 48)
	z01.PrintRune('\n')
}
