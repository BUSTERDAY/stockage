package main

import (
	"fmt"
)

func StrRev(s string) string {
	var str string
	for i := len(s) - 1; i > 0; i-- {
		str += string(s[i])
	}
	str += string(s[0])
	return str
}

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
