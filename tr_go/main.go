package main

import "fmt"

func CheckAge(age int) bool {
	if age >= 18 {
		fmt.Print("Tu peux entrer")
		return true
	}
	fmt.Print("Sortez !")
	return false
}

func CheckIdentity(age int, sexe rune) {
	if CheckAge(age) {
		if sexe == 70 || sexe == 102 {
			fmt.Println(", c'est 10€")
		} else if sexe == 72 || sexe == 104 {
			fmt.Println(", c'est 15€")
		}
	}
}
