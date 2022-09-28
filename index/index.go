package main

func Index(s string, toFind string) int {
	var comp = []rune(toFind)
	var count int
	for i, letter := range s {
		if letter == comp[0] {
			return i
		}
	}
	return count
}

func main() {
	println(Index("Hello how are u z", "u z"))
}
