package main

import "fmt"

func main() {

	c := make(chan int, 2)
	cr := make(<-chan int) //recieve
	cs := make(chan<- int) // send

	fmt.Println("-----------")
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

	// spesific to general doesnt assign
	// c = cr
	// c = cs

	// this one works
	cr = c
	cs = c

}
