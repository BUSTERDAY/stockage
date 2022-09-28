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

func intToTab(n int) []int {
	var count int
	var tmp int
	var mod int = 1
	var tab []int
	for tmp = n; tmp != 0; tmp /= 10 {
		count++
		mod *= 10
	}
	for i := 0; i != count; i++ {
		tmp = n
		tmp %= mod
		for ; tmp > 9; tmp /= 10 {
		}
		tab[i] = tmp
		mod /= 10
	}
	return tab
}

func PrintNbrInOrder(n int) {
	//var i int
	//var j int
	var tab []int = intToTab(n)
	println(tab)
}

func main() {
	PrintNbrInOrder(12334980)
}
