package main

import "fmt"

const a = 42
const b = 42.354
const c = "James Bond"

const (
	d = 42
	e = 42345
	f = "James Bond"
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", d)

}
