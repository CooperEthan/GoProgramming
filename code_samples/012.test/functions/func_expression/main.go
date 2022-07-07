package main

import "fmt"

func main() {
	fmt.Println("Hello")

	f := func() {
		fmt.Println("myfirst func expression")
	}
	f()

	f1 := func(x int) {
		fmt.Println("my second func", x)
	}
	f1(42)

}
