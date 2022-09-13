package main

import "fmt"

func CheckAge(age int) bool {
	if age < 18 {
		return false
	} else {
		return true
	}
}

func CheckIdentity(age int, sexe string) {
	if !CheckAge(age) {
		fmt.Println("Sortez")
	} else {
		if sexe == "Femme" {
			fmt.Println("Tu peux entrer, c'est 10€")
		} else {
			fmt.Println("Tu peux entrer, c'est 15€")
		}
	}
}

func main() {
	CheckIdentity(18, "Femme")
	CheckIdentity(24, "Homme")
	CheckIdentity(15, "Femme")
}
