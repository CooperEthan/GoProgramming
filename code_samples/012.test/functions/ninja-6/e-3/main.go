package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func main() {

	p := person {
		first: "James",
		last: "Bond",
		age: 20,
	}
	p.speak()

}

func (p person) speak() {
	fmt.Println("I am",p.first,p.last," and I am ",p.age, " years old.")
}