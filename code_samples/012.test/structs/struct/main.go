package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   40,
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   30,
	}

	fmt.Println(p1.first, p2.last, p1.age)
	fmt.Println(p2.age)

}
