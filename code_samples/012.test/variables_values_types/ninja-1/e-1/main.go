package main

import "fmt"

var x int =42
var y string = "James Bond"
var z bool = true

var t int
var v string

type ethan int
var bx ethan

func main(){

	bx = 100
	fmt.Println(bx)

	x = int(bx)
	fmt.Printf("%T\t", x)
	fmt.Println("\t", x)
	fmt.Printf("%T", bx)

	fmt.Println("====================")
	bx = ethan(x)
	fmt.Printf("%T\n", x)

	fmt.Println("==========")
	z := 50 
	fmt.Println(z)

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println("===================")

	fmt.Println(t)
	fmt.Println(v)

	fmt.Println("===================")
	s := fmt.Sprintf("%v\t%v\t%v\t",x,y,z)
	fmt.Println(s)
}