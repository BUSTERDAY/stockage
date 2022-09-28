package main

func Capitalize(s string) string {
	var str []rune = []rune(s)
	var i int
	if str[0] >= 97 && str[0] <= 122 {
		str[0] -= 32
	}
	for i = 1; i != len(s); i++ {
		if str[i] >= 32 && str[i] <= 47 && str[i+1] >= 97 && str[i+1] <= 122 && i != len(s)-1 {
			str[i+1] -= 32
		} else if str[i] >= 64 && str[i] <= 90 && !(str[i] >= 32) && !(str[i] <= 47) && i != len(s)-1 {
			str[i] += 32
		}
	}
	return string(str)
}

func main() {
	println(Capitalize("Hello! How are you? how+are!things+4you?"))
}
