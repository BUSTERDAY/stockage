package main

import "fmt"

func AppendRange(min, max int) []int {
	var liste []int
	for i := min; i < max; i++ {
		liste = append(liste, i)
	}
	return liste
}

func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}
