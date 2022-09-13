package main

import "fmt"

func FindNextPrime(nb int) int {
	for i := nb; i > 0; i++ {
		if isprime(i) {
			return i
		}
	}
	return 0
}

func isprime(nb int) bool {
	if nb == 0 {
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
	fmt.Println(FindNextPrime(5))
	fmt.Println(FindNextPrime(4))
}
