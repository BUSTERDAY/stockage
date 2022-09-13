package main

import "fmt"

func MakeRange(min, max int) []int {
	if min > max {
		var vide []int
		return vide
	}
	var liste = make([]int, max-min)
	var nb int = min
	if min < max {
		for i := 0; i != max-min; i++ {
			liste[i] = nb
			nb++
		}
	}
	return liste
}

func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}
