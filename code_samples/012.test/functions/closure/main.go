package main

import (
	"fmt"
)

func main() {

	var x int

	fmt.Println(x)
	x++
	fmt.Println(x)
	{
		y := 1
		fmt.Println(y)
	}
	foo()
	fmt.Println(x)

	fmt.Println("=================================")
	a := incrementer()
	b := incrementer()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())

}

func foo() {
	fmt.Println("hello")
}

func incrementer() func() int {
	var z int
	return func() int {
		z++
		return z
	}
}
