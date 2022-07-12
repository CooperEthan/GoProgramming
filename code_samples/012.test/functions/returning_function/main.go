package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)
	fmt.Println(foo())

	x := bar()
	fmt.Println(x())

	fmt.Println(bar()())

	// i := x()

	// fmt.Println(i)
	// fmt.Printf("%T", i)
}

func foo() string {
	s := "Hello World"
	return s
}

// a func returns a func that returns int
func bar() func() int {
	return func() int {
		return 451
	}
}
