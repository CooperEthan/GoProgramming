package main

import "fmt"


type person struct {
	first string
	age int
}
type animal struct{
	name string
}
type human interface {
	speak ()
}

func main() {
	p1 := person {
		first: "James",
		age: 20,
	}
	saySomething(&p1)
	p1.speak()
	
	a := animal {
		name: "lion",
	}
	a.speak()
	saySomething(&a)
}

func (p *person) speak() {
	fmt.Println(p.first,p.age)
}
func (a *animal) speak() {
	fmt.Println(a.name)
}
func saySomething(h human) {
	h.speak()
}