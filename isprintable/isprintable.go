package main

func IsPrintable(s string) bool {
	var str = []byte(s)
	for i := 0; i != len(s); {
		if str[i] >= 32 && str[i] <= 126 {
			i++
		} else {
			return false
		}
	}
	return true
}

func main() {
	println(IsPrintable("Hello"))
	println(IsPrintable("Hello\n"))
}
