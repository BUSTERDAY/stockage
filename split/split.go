package main

import (
	"fmt"
)

func Split(s, sep string) []string {
	var resultat []string
	var temp string
	for i, j := range s {
		temp += string(j)
		for _, a := range sep {

		}
		resultat = append(resultat, temp)
		temp = ""
	}
	resultat = append(resultat, temp)
	return resultat
}

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}
