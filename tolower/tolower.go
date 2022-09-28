package main

func ToLower(s string) string {
	var str = []rune(s)
	var i int
	for i = 0; i != len(s); i++ {
		if str[i] >= 65 && str[i] <= 90 {
			str[i] += 32
		}
	}
	return string(str)
}

func main() {
	println(ToLower("testTE"))
}
