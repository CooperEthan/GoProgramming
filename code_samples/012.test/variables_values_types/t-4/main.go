package main

import "fmt"

var y string
var x int

func main() {

	y = "I am Yakut"

	x = 5

	fmt.Println("printing the value of y", y, "ending")
	fmt.Printf("%T\n", y)
	fmt.Println(x)

	fmt.Printf("%b\n", x)
	fmt.Printf("%#x\t%b\t%x\n", x, x, x)

	s := fmt.Sprintf("%#x\t%b\t%x", x, x, x)

	fmt.Println(s)

	fmt.Printf("%v", y)

}
