package main

import "fmt"

func main() {

	a := 42
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
	b := &a
	fmt.Println(b)
	fmt.Println(*b)  // * gives you the value stored at an address when you have the address
	fmt.Println(*&a) // gives you the value

}
