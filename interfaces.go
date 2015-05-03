package main

import "fmt"

// convention: "er" a type to indicate that it is an interface
// our interface has a single function returning a string
type Greater interface{
	SayHello() string
}

// Person is a struct with a name
type Person struct {
	name string
}

type Dog struct {
	name string
}
// Person is implementing SayHello interface
func (p Person) SayHello() string {
	return "Hello world!"
}

func (d Dog) SayHello() string {
	return "Grrr Grrrr!"
}

func main() {
	p := Person{name:"pitt"}
	d := Dog{name:"bill"}

	fmt.Println(p)
	fmt.Println(p.SayHello())

	// you can cast d to be a Greater type
	g := Greater(d)
	fmt.Println(d)
	// then call the SayHello function on it
	fmt.Println(g.SayHello())
}
