package main

func IsNumeric(s string) bool {
	var str = []byte(s)
	for i := 0; i != len(s); {
		if str[i] >= 48 && str[i] <= 57 {
			i++
		} else {
			return false
		}
	}
	return true
}

func main() {
	println(IsNumeric("010203"))
	println(IsNumeric("01,02,03"))
}
