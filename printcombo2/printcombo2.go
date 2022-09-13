package main

import "github.com/01-edu/z01"

func PrintComb2() {
	for a := 48; a < 58; a++ {
		for b := 48; b < 57; b++ {
			for c := 48; c < 58; c++ {
				for d := 48; d < 58; d++ {
					if a*10+b < c*10+d {
						z01.PrintRune(rune(a))
						z01.PrintRune(rune(b))
						z01.PrintRune(rune(','))
						z01.PrintRune(rune(c))
						z01.PrintRune(rune(d))
						if a == '9' && b == '8' && c == '9' && d == '9' {
							return
						} else {
							z01.PrintRune(rune(','))
							z01.PrintRune(rune(' '))
						}
					}
				}
			}
		}
	}
}

func main() {
	PrintComb2()
}
