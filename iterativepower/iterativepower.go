package main

import "fmt"

func IterativePower(nb int, power int) int {
	var resultat int = 1
	for i := power; i > 0; i-- {
		resultat *= nb
	}
	return resultat
}

func main() {
	fmt.Println(IterativePower(4, 3))
}
