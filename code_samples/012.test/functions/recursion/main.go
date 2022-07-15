package main

import (
	"fmt"
)

func main() {
	n := factorial(4)
	fmt.Println(n)

	x := loop(4)
	fmt.Println(x)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loop(n int) int {
	a := 1
	for ; n > 0; n-- {
		a *= n
	}
	return a
}
