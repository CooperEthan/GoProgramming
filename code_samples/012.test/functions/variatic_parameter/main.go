package main

import "fmt"


func main() {
// Unfurling slice
	xi := []int{2,3,4,5,6}
	x := foo(xi...)

	// foo(2,3,4,5,6)
	// x := foo(2,3,4,5,6)
	fmt.Println(x)

}

func foo(x ...int) int{
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	fmt.Println("Capacity is: ",cap(x))
	fmt.Println("Length is: ",len(x))

	sum := 0
	for i, v := range x {
		sum += v 
		fmt.Println("for index position, ", i, "we are adding, ", v, "to the total now")
	}
	fmt.Println("total is ", sum)
	return sum 
}