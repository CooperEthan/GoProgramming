package main

import (
	"fmt"

	"example.com/v2/writing_documentation/dog"
)

type canine struct {
	name string
	age  int
}

func main() {

	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)

}

// func Sum(xi ...int) int {
// 	sum := 0
// 	for _, v := range xi {
// 		sum += v
// 	}
// 	fmt.Println(sum)
// 	return sum

// }
