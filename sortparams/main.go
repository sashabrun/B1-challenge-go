package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	cassecouille := os.Args[1:]
	for i := len(cassecouille) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if cassecouille[j] > cassecouille[j+1] {
				cassecouille[j], cassecouille[j+1] = cassecouille[j+1], cassecouille[j]
			}
		}
	}
	if len(cassecouille) > 0 {
		for _, cassecouille := range os.Args[1:] {
			for i := 0; i < len(cassecouille); i++ {
				z01.PrintRune(rune(cassecouille[i]))
			}
			z01.PrintRune('\n')
		}
	}
}
