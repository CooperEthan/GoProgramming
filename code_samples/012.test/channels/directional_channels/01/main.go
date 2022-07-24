package main

import "fmt"

func main() {

	c := make(chan int, 2)
	cr := make(<-chan int) //recieve
	cs := make(chan<- int) // send

	// c <- 42
	// c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("-----------")
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)
}
