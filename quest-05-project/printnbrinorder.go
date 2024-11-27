package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	variables := []int{}
	for n > 0 {
		variables = append(variables, n%10)
		n = n / 10
	}

	for i := 0; i < len(variables); i++ {
		for j := 0; j < len(variables)-i-1; j++ {
			if variables[j] > variables[j+1] {
				count := variables[j]
				variables[j] = variables[j+1]
				variables[j+1] = count
			}
		}
	}
	for _, val := range variables {
		z01.PrintRune(rune(val + '0'))
	}
}
