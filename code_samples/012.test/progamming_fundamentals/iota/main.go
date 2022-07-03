package main

import "fmt"



const (
	a = iota
	b = iota
	c = iota
	k
	l
	m
)

const (
	d = iota
	e 
	f 
)

func main () {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

}