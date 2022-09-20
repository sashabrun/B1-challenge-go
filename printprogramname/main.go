package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	titredefdp := []rune(os.Args[0])[1:]
	for i := 0; i < len(titredefdp); i++ {
		z01.PrintRune(rune(titredefdp[i]))
	}
}
