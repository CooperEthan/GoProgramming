package main

import "fmt"

func main() {
	func() {
		x := 5
		fmt.Println(x)
	}()

	f := foo()
	g := foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(g())
	fmt.Println(g())
}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
