package main

import "fmt"

var resultat int = 1

func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if power == 1 {
		resultat = nb
	}
	if power > 0 {
		return nb * RecursivePower(nb, power-1)
	}
	return 0
}

func main() {
	fmt.Println(RecursivePower(4, 3))
}
