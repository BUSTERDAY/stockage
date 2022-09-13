package main

import "fmt"

func PrintWordsTables(a []string) {
	for _, i := range a {
		fmt.Println(i)
	}
}

func SplitWhiteSpaces(s string) []string {
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
	a := SplitWhiteSpaces("Hello how are you?")
	PrintWordsTables(a)
}
