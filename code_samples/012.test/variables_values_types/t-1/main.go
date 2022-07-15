package main

import "fmt"

func main() {
	fmt.Println("Hello ethan, this is going to be awesome")

	n, _ := fmt.Println("Hello", 42, true)
	fmt.Println(n)

	fmt.Println("more")

	for i := 0; i < 9; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	foo()
}

func foo() {
	fmt.Println("now exiting!")
}
