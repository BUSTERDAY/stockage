package main

import (
	"fmt"
)

/*func Split(s, sep string) []string {
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
}*/

func Split(s, sep string) []string {
	var tmp string
	var i int
	var count int
	for _, letter := range s {
		if letter == ' ' || letter == '\n' || letter == '\t' {
			count++
		}
	}
	res := make([]string, count+1)
	for j := 0; j != len(s); j++ {
		if s[j] == ' ' || s[j] == '\n' || s[j] == '\t' {
			res[i] = tmp
			tmp = ""
			i++
			j++
		}
		tmp += string(s[j])
	}
	res[i] = tmp
	return res
}

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}
