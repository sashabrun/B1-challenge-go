package main

import (
	"os"
)

func main() {
	cassecouille := os.Args[1:]
	n := len(cassecouille)
	b := false
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if cassecouille[j] > cassecouille[j+1] {
				cassecouille[j], cassecouille[j+1] = cassecouille[j+1], cassecouille[j]
				b = true
			}
		}
		if !b {
			print(cassecouille[i])
		}
	}
}
