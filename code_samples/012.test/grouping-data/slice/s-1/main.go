package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Println(len(x))

	fmt.Println("number 2 is: ", x[1])
	for _, y := range x {
		fmt.Println(y)
	}
}
