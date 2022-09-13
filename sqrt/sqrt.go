package main

import "fmt"

func Sqrt(nb int) int {

	for c := 1; c < nb; c++ {
		if RecursivePower(c, 2) == nb {
			return c
		}
	}
	return 0
}

func RecursivePower(nb int, power int) int {

	if power == 1 {
		return nb
	}
	if power > 0 {
		return nb * RecursivePower(nb, power-1)

	}
	return 0
}

func main() {
	arg := Sqrt(9)
	fmt.Printf("%v", arg)
}
