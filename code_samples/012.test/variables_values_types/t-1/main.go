package main

import "fmt"

func main() {
	fmt.Println("Hello ethan")

	foo()
	fmt.Println("more")

	for i := 0; i<9; i++ {
		if i %2 == 0{
		fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("I am foo")
}

func bar(){
	fmt.Println("now exiting")
}