package main

import "fmt"

func main() {

	var x [5]int

	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
	x[4] = 5

	for i, xi := range x {
		fmt.Println(i, xi)
	}
	// =============================================

	y := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(y[:5])
	fmt.Println(y[5:])
	fmt.Println(y[2:7])
	fmt.Println(y[1:6])

	//==============================================

}
