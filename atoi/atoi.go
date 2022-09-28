package main

import (
	"fmt"
)

func Atoi(s string) int {
	var result int
	var negative bool = false
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			return 0
		}
		if s[i] > 48 && s[i] <= 57 {
			result *= 10
			result += (int(s[i]) - 48)
		}
		if s[i] == '-' {
			negative = true
		}
		if s[i] == '-' && s[i+1] == '-' {
			return 0
		} else if s[i] == '+' && s[i+1] == '+' {
			return 0
		}
	}
	if negative {
		result *= -1
		return result
	} else {
		return result
	}
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}
