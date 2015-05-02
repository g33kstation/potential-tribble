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

func beyondHello() {
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
	return x + y, x * y

}

func learnTypes() {
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

	// byte is an alias for uint8
	n := byte('\n')

	// array are size fixed
	var a4 [4]int           // array size 4, int items 0, 0, 0, 0
	a3 := [...]int{3, 1, 5} // array size 3, int items 3, 1, 5

	// slices are size not fixed
	s3 := []int{4, 5, 9} // 3 ints slice
	s4 := make([]int, 4) // 4 ints slice initialized with zeros
	var d2 [][]float64
	bs := []byte("a slice")

	// add values to slice
	s3 = append(s3, 4, 5, 6)
	s3 = append(s3, []int{7, 8, 9}...)
	fmt.Println(s3)

	// maps are like python dictionaries
	m := map[string]int{"score": 5}
	m["distance"] = 5
	fmt.Println(m)

	// declared & not used variables are errors
	// to ignore those unused values
	_, _, _, _, _, _, _, _, _, _, _, _, _ = str, s2, g, f, c, u, pi, n, a3, a4, bs, d2, s4
	learnFlowControl()
}

func learnFlowControl() {
	if true {
		fmt.Println("bim")
	}
	if false {
		// if
	}else{
		// else
	}
	// switch case
	x := 42.0
	switch x {
	case 0:
	case 1:
		fmt.Println("0 or 1")
	case 42:
		fmt.Println("42")
	case 43:
		fmt.Println("43")
	}

	// for loop
	for x:=0; x<3; x++ {
		fmt.Println("i: ", x)
	}

	// k,v will go through the different k,v pairs of the map
	for k, v := range map[string]int{"une":1, "deux":2, "trois":3}{
		fmt.Printf("key=%s, value=%d\n", k, v)
	}

	//
}
