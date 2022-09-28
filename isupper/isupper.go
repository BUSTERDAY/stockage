package main

func IsUpper(s string) bool {
	var str = []byte(s)
	for i := 0; i != len(s); {
		if str[i] >= 65 && str[i] <= 90 {
			i++
		} else {
			return false
		}
	}
	return true
}

func main() {
	println(IsUpper("HELLO"))
	println(IsUpper("HELLO!"))
}
