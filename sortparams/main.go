package main

import "os"

func main() {
	cassecouille := []rune(os.Args[0])[2:]
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
			return
		}
	}
}
