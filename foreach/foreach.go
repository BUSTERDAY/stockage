package main

import "github.com/01-edu/z01"

func ForEach(f func(int), a []int) {
	for _, arg := range a {
		f(arg)
	}
}

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
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
}
