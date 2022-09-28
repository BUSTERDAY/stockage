package main

import "fmt"

func doop(a int, signe string, b int) int {
	var resultat int
	switch signe {
	case "+":
		resultat = a + b
	case "-":
		resultat = a - b
	case "*":
		resultat = a * b
	case "/":
		resultat = a / b
	case "%":
		resultat = a % b
	}
	return resultat
}

func main() {
	fmt.Println(doop(2, "+", 4))
	fmt.Println(doop(2, "-", 4))
	fmt.Println(doop(2, "*", 4))
	fmt.Println(doop(2, "/", 4))
	fmt.Println(doop(2, "%", 4))
}
