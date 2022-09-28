package main

import (
	"fmt"
)

func BasicAtoi2(s string) int {
	var result int
	for _, i := range s {
		if i == ' ' {
			return 0
		}
		if i >= 97 && i <= 122 {
			return 0
		}
		if i > 48 && i <= 57 {
			result *= 10
			result += (int(i) - 48)
		}
	}

	return result
}

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}
