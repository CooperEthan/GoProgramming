package main

import "fmt"


func main () {

	// x := 83 / 40
	// y := 83 % 40
	// fmt.Printf("x is: %d\ny is: %d\t", x, y)

	x := 0
	for {
		x++
		if x >10 {
			break
		}
		if x%2 != 0{
			continue
		}
		fmt.Println(x)
	}


}