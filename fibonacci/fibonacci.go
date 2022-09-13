package main

import "fmt"

func fibonacci(index int) int {
	if index < 0 {
		return -1
	} else if index <= 1 {
		return index
	} else {
		return (fibonacci(index-1) + fibonacci(index-2))
	}
}

func main() {
	arg1 := 4
	fmt.Println(fibonacci(arg1))
}
