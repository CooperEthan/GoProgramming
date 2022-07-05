package main

import (
	"fmt"
)


type person struct {
	// first string
	// last string
	first,last string
	age int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {

	sa1 := secretAgent {
		person: person{
			first: "James",
			last: "Bond",
			age: 30,
		},
		ltk: true,
		}

	fmt.Println(sa1)

	x := 42
	y := 43
	fmt.Println(x,y)
	x,y = y, x
	fmt.Println(x,y)

}