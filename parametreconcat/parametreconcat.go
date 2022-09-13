package main

import "fmt"

func ConcatParams(args []string) string {
	var resultat string
	for _, i := range args {
		resultat += i + "\n"
	}
	return resultat
}

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}
