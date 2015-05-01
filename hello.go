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
}

// arguments x, y of type integers
// this function returns 2 integers, sum & prod
func learnMultiple(x, y int) (sum, prod int) {
	// two values returned
	return x+y, x*y

}
