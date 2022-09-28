package main

func IsLower(s string) bool {
	var str = []byte(s)
	for i := 0; i != len(s); {
		if str[i] >= 97 && str[i] <= 122 {
			i++
		} else {
			return false
		}
	}
	return true
}

func main() {
	println(IsLower("hello"))
	println(IsLower("hello!"))
}
