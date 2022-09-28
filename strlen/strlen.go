package main

import (
	"fmt"
)

func StrLen(s string) int {
	var compteur int
	for range s {
		compteur += 1
	}
	return compteur
}

func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}
