package piscine

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	even := 0
	odd := 0
	for x := 1; x <= nbr; x++ {
		if x%2 == 0 {
			even = even + x
		} else {
			odd = odd + x
		}
	}
	if isEven(nbr) {
		return true
	} else {
		return false
	}
}

func main() {
	y := os.Args[1:]
	if isEven(len(y)) {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
