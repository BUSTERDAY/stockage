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

func TrimAtoi(s string) int {
	var str = []byte(s)
	var res int = 0
	var isNeg bool = false
	for i := 0; i != len(s); i++ {
		if str[i] == 45 && !(res > 0) {
			isNeg = true
		}
		if IsNumeric(string(str[i])) {
			if res != 0 {
				res *= 10
			}
			res += int(str[i] - 48)
		}
	}
	if isNeg == true {
		res *= -1
	}
	return res
}

func main() {
	println(TrimAtoi("12345"))
	println(TrimAtoi("str123ing45"))
	println(TrimAtoi("012 345"))
	println(TrimAtoi("Hello World!"))
	println(TrimAtoi("sd+x1fa2W3s4"))
	println(TrimAtoi("sd-x1fa2W3s4"))
	println(TrimAtoi("sdx1-fa2W3s4"))
	println(TrimAtoi("sdx1+fa2W3s4"))
}
