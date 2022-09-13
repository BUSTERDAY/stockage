package main

func IsNegative(nb int) {
	if nb < 0 {
		println("T")
	} else {
		println("F")
	}
}

func main() {
	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)
}
