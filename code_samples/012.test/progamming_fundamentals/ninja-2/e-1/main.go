package main

import "fmt"


func main() {
	x := 42

	fmt.Printf("%b\n", x)
	fmt.Printf("%d\n", x)
	fmt.Printf("%#x\n", x)

	fmt.Printf("%d\t%b\t%#x", x, x, x)
}
