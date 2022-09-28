package main

func IsAlpha(s string) bool {
	var str = []byte(s)
	for i := 0; i != len(s); {
		if str[i] >= 97 && str[i] <= 122 || str[i] >= 65 && str[i] <= 90 || str[i] >= 48 && str[i] <= 57 {
			i++
		} else {
			return false
		}
	}
	return true
}

func main() {
	println(IsAlpha("Hello! How are you?"))
	println(IsAlpha("HelloHowareyou"))
	println(IsAlpha("What's this 4?"))
	println(IsAlpha("Whatsthis4"))
}
