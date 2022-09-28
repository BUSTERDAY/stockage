package main

import (
	"fmt"
)

func Map(f func(int) bool, a []int) []bool {
	var resultat []bool
	for _, arg := range a {
		resultat = append(resultat, f(arg))
	}
	return resultat
}

func IsPrime(nb int) bool {
	if nb == 0 {
		return false
	} else if nb == 1 {
		return false
	}

	for i := nb - 1; i > 1; i-- {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(IsPrime, a)
	fmt.Println(result)
}
