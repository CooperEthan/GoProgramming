package main

import (
	"fmt"
)

var x int
var g func()

func main() {

	a := func ()  {
		fmt.Println("I am ananymous func")
	}
	a()
	fmt.Printf("%T",a)

	g = a
	fmt.Printf("%T",g)

	defer func (x int)  {
		fmt.Println(x)
	}(5)
	defer func (y int)  {
		for i := 0; i < y; i++ {
			fmt.Println(i)
		}
	}(5)

	// fmt.Println("===================")

	x := foo(5)
	fmt.Println("I will be the first one the called in main -->",x)
}

func foo(x int) int {
	sum :=0
	for i := 0; i < x; i++ {
		sum +=i
	}
	return sum
}