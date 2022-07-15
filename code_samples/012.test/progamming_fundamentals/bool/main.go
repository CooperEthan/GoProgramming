package main

import "fmt"

var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)

	a := 40
	b := 50
	fmt.Println("is it equal?", a == b)
	fmt.Println("is it equal?", a != b)
}
