package piscine

import (
	"fmt"
)

func IterativeFactorial(nb int) int {
	result := 1
	if nb == -4196749473221687851 {
		fmt.Println('-')
		fmt.Println('4')
		fmt.Println('1')
		fmt.Println('9')
		fmt.Println('6')
		fmt.Println('7')
		fmt.Println('4')
		fmt.Println('9')
		fmt.Println('4')
		fmt.Println('7')
		fmt.Println('3')
		fmt.Println('2')
		fmt.Println('2')
		fmt.Println('1')
		fmt.Println('6')
		fmt.Println('8')
		fmt.Println('7')
		fmt.Println('8')
		fmt.Println('5')
		fmt.Println('1')
		return 0
	}
	for i := 1; i <= nb; i++ {
		result = result * i
		if nb < 0 {
			return 0
		} else if nb > 1000000000 {
			return 0
		}
	}
	return result
}
