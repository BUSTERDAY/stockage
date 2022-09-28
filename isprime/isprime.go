package main

import "fmt"

func IsPrime(nb int) bool {
	if nb == 0 {
		return false
	} else if nb == 1 {
		return false
	}

	for i := nb - 1; i > 1; i-- {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(4))
}
