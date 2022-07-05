package main

import "fmt"


func main() {
	var x [5]int
	fmt.Println(x)

	x [3] = 43
	fmt.Println(x[3])
	fmt.Println(len(x))
}