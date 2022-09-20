package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 0 {
		for i := len(os.Args) - 1; i >= 1; i-- {
			for _, char := range os.Args[i] {
				z01.PrintRune(rune(char))
			}
			z01.PrintRune('\n')
		}
	}
}
