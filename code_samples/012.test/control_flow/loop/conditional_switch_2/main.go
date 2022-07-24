package main

import (
	"fmt"
)

func main() {
	// 	n := false
	// 	switch n {
	// 	case false:
	// 		fmt.Println("this should not print")
	// 	case (2==4):
	// 		fmt.Println("this should not print2")
	// 	case (3==3):
	// 		fmt.Println("prints")
	// 		fallthrough
	// 	case (4==4):
	// 		fmt.Println("this should not print?")
	// 		fallthrough
	// 	case (4==4):
	// 		fmt.Println("this should not print?")
	// 		fallthrough
	// 	default:
	// 		fmt.Println("default case")
	// }

	switch {
	case false:
		fmt.Println("this should not print")
	case true:
		fmt.Println("this should print")
	}

}
