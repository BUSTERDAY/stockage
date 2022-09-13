package main

import "fmt"

var resultat int = 1

func RecursiveFactorial(nbr int) (return_type int) {
	if nbr == 0 {
		return resultat
	} else {
		resultat *= nbr
		return RecursiveFactorial(nbr - 1)
	}
}

func main() {
	arg := 4
	fmt.Println(RecursiveFactorial(arg))
}
