package main

import "fmt"


type person struct {
	first string
	last string
}
type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println(s.first,s.last,s.ltk,s.person)
}

func main() {

	sa1 := secretAgent {
		person: person{
			first: "James",
			last: "Bond",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa2 := secretAgent {
		person: person{
			first: "James",
			last: "Bond",
		},
		ltk: true,
	}
	fmt.Println(sa2)

	sa1.speak()
	sa2.speak()

}