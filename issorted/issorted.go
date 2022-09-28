package main

import "fmt"

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}

	result1 := IsSorted(ecart, a1)
	result2 := IsSorted(ecart, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if f(int(a[i]), int(a[i+1])) < 1 {
			return false
		}
	}
	return true
}

func ecart(a, b int) int {
	return b - a
}
