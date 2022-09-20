package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 0 {
		for _, char := range os.Args[1:] {
			for i := len(char) - 1; i >= 0; i-- {
				z01.PrintRune(rune(char[i]))
			}
			z01.PrintRune('\n')
		}
	}
}
