package main

func AlphaCount(s string) int {
	var str = []rune(s)
	var count int
	for i := 0; i != len(s); i++ {
		if str[i] >= 65 && str[i] <= 90 || str[i] >= 97 && str[i] <= 122 {
			count += 1
		}
	}
	return count
}

func main() {
	println(AlphaCount("a B c D eF"))
}
