package main

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var count int
	var tmp int
	var mod int = 1
	if n < 0 {
		n *= -1
		z01.PrintRune('-')
	}
	for tmp = n; tmp != 0; tmp /= 10 {
		count++
		mod *= 10
	}
	for ; count != 0; count-- {
		tmp = n
		tmp %= mod
		for ; tmp > 9; tmp /= 10 {
		}
		z01.PrintRune(rune(tmp + 48))
		mod /= 10
	}
}

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}
