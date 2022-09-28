package main

import (
	"fmt"
)

func IsNumeric(s string) bool {
	var str = []byte(s)
	for i := 0; i != len(s); {
		if str[i] >= 48 && str[i] <= 57 {
			i++
		} else {
			return false
		}
	}
	return true
}

func CountIf(f func(string) bool, tab []string) int {
	var resultat int
	for i, arg := range tab {
		if f(arg) == true {
			resultat += 1
		}
		i++
	}
	return resultat
}

func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This", "1", "is", "4", "you"}
	answer1 := CountIf(IsNumeric, tab1)
	answer2 := CountIf(IsNumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}
