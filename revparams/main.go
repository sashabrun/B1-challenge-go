package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 0 {
		for i := len(os.Args[1:]) - 1; i >= 0; i-- {
			for _, char := range os.Args[1:] {
				z01.PrintRune(rune(char[i]))
			}
			z01.PrintRune('\n')
		}
	}
}
