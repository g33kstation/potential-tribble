package main

import "fmt"

// a function which takes a string as argument
// and return a function as result
func sentenceFactory(s string) func(b, a string) string {
	// this result, is a function taking 2 strings as arguments
	return func(b, a string) string {
		// and returning a string
		return fmt.Sprintf("%s %s %s", b, s, a)
	}
}

func main() {
	x := sentenceFactory("patrick")
	// decorating x
	fmt.Println(x("the awesome ", "the amazing"))
}
