package main

func ToUpper(s string) string {
	var str = []rune(s)
	var i int
	for i = 0; i != len(s); i++ {
		if str[i] >= 97 && str[i] <= 122 {
			str[i] -= 32
		}
	}
	return string(str)
}

func main() {
	println(ToUpper("testTE"))
}
