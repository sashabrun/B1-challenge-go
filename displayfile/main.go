package main

import (
	"fmt"
	"os"
)

func main() {
	connard := os.Args
	if len(connard) < 2 && len(connard) > 0 {
		fmt.Println("File name missing")
	} else if len(connard) > 2 {
		fmt.Println("Too many arguments")
	} else if len(connard) == 2 {
		connard, _ := os.ReadFile("quest8.txt")
		fmt.Print(string(connard))
	}
}
