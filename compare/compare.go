package main

func Compare(a, b string) int {
	if len(a) != len(b) {
		return 1
	}
	var str1 = []rune(a)
	var str2 = []rune(b)
	for i := 0; i != len(a)-1; {
		if str1[i] == str2[i] {
			i++
		} else {
			return 1
		}
	}
	return 0
}

func main() {
	println(Compare("Hello", "Helli"))
}
