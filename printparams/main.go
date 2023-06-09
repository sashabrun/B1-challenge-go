package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 0 {
		for _, char := range os.Args[1:] {
			for i := 0; i < len(char); i++ {
				z01.PrintRune(rune(char[i]))
			}
			z01.PrintRune('\n')
		}
	}
}
