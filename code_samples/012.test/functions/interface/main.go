package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

type human interface {
	speak()
}

func bar(h human) {
	fmt.Println("I called human", h)
}

type hotdog int

func main() {

	sa1 := secretAgent{
		person: person{
			first: "Matt",
			last:  "Damon",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			first: "Kevin",
			last:  "Kostner",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr",
		last:  "No",
	}

	// *** ananymous fun *** ================
	defer func() {
		fmt.Println("Ananymous func")
	}()
	defer func(x int) {
		fmt.Println("Ananymous func with argument", x)
	}(42)
	//=========================================================
	sa1.speak()
	sa2.speak()
	p1.speak()

	bar(sa1)
	bar(sa2)
	bar(p1)

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T", x)
	y := int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)

}

func (s secretAgent) speak() {
	fmt.Println(s.first, s.last, s.ltk)
}

func (p person) speak() {
	fmt.Println(p.first, p.last)
}
