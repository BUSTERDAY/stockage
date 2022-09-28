package main

func Join(elems []string, sep string) string {
	var res string
	for i, elem := range elems {
		if i != 0 {
			res += sep
		}
		res += elem
	}
	return res
}

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	println(Join(toConcat, ":"))
}
