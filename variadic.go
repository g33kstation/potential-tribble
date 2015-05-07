package main

import "fmt"

// this func takes zero or more int as arguments
func add(args ...int) int {
	total := 0
	// ignore the index arg of current item
	for _, v := range args {
		total += v
	}
	return total
}

func index(args ...int) int {
	total := 0
	// k is the index of the current item
	for k, v := range args {
		// any treatment using k
		total = total - k + v
	}
	return total
}

func main() {
	// we can pass as many ints as we want
	fmt.Println(add(1,2,3))
	fmt.Println(index(3,1,2))

	// another way of calling
	xs := []int{1,2,3,4}
	// unpacking?!
	fmt.Println(add(xs...))
	fmt.Println(index(xs...))
}
