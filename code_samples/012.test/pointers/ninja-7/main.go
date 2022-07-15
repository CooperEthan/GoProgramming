package main

import "fmt"

type person struct {
	name string
}

func changeMe(p *person) {
	p.name = "MisMoneypenny"
	(*p).name = "James Bond"

}

func main() {

	a := 42
	fmt.Println("value is=", a)
	fmt.Println("Adress of the value=", &a)
	fmt.Println("====================")

	p1 := person{
		name: "James Bond",
	}
	changeMe(&p1)
	fmt.Println(*&p1)
	// changeMe(&p1)
	fmt.Println(&p1)
}
