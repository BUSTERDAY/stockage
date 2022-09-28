package main

func BasicJoin(elems []string) string {
	var res string
	for _, elem := range elems {
		res += elem
	}
	return res
}

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	println(BasicJoin(elems))
}
