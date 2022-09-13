package main

import "fmt"

func IterativeFactorial(nbr int) (return_type int) {
	var resultat int = 1
	for i := 1; i <= nbr; i++ {
		resultat = resultat * i
	}
	return resultat
}

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}
