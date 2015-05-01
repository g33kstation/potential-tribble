// package declared as executable, not lib
package main

// an import
import "fmt"

// entry point
func main() {
	// call Printf function from fmt package
	fmt.Printf("hello, world\n")

	// call another function
	beyondHello()

}

func beyondHello(){
	// variable should be declared before being used
	var x int
	// assignation
	x = 3

	// type inference, declaration, assignation
	y := 4

	// a function returning 2 values
	sum, prod := learnMultiple(x, y)
	fmt.Println("sum:", sum, "prod:", prod)

	// another function call
	learnTypes()
}

// arguments x, y of type integers
// this function returns 2 integers, sum & prod
func learnMultiple(x, y int) (sum, prod int) {
	// two values returned
	return x+y, x*y

}

func learnTypes(){
	// string inference
	str := "Learn go"
	s2 := `string with
newline`
	// go source code use utf8 charset
	// rune type, for unicode char
	g := 'Σ'

	// float
	f := 3.14195
	// complex128 type, considered as float64 by compiler
	c := 3 + 4i
	// u declared as unsigned int
	var u uint = 7
	var pi float32 = 22. / 7

	// declared & not used variables are errors
	// to ignore those unused values
	_, _, _, _, _, _, _ = str, s2, g, f, c, u, pi
}
