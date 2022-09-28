package main

import (
	"fmt"
)

func BasicAtoi(s string) int {
	var result int
	for _, i := range s {
		if i > 48 && i <= 57 {
			result *= 10
			result += (int(i) - 48)
		}
	}

	return result
}

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}
