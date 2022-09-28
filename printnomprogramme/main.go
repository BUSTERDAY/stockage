package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	for _, i := range arguments {
		for _, j := range i {
			z01.PrintRune(j)
		}
	}

}
