package main

import "fmt"

// type person struct{
// 	first string
// 	last string
// 	age int
// }

func main() {

	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   40,
	}

	fmt.Println(p1)

}
