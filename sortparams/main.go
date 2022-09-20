package main

import (
	"github.com/01-edu/z01"
	"os"
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
	for _, nom := range cassecouille {
		print(nom)
		z01.PrintRune('\n')
	}
}
