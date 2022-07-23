package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5}
	// fmt.Println(x[1])
	// fmt.Println(x[1:])
	// fmt.Println(x[1:3])
	// fmt.Println(x[:2])

	// for i, v := range x {
	// 	fmt.Println(i, v,)
	// }

	x = append(x, 10, 100, 1000)

	// for i := 0; i < 8; i++ {
	// 	fmt.Println(i, x[i])
	// }

	y := []int{123, 345, 456, 567}
	x = append(x, y...)
	// x = append(y, 35,40,50 )

	fmt.Println("\n", x)

	x = append(x[:2], x[4:]...)

	fmt.Println(x)

}
