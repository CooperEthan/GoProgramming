package main

import (
	"fmt"
	"math"
)

type square struct {
	a int
}
type circle struct {
	r float64
}
type shape interface {
	area()
}
// func info(s shape){
// 	fmt.Println(s.area())
// }

func main() {
	r := circle {
		r: 42,
	}
	s := square {
		a: 42,
	}
	r.area()
	s.area()
}

func(s square) area() {
	b := (s.a)*(s.a)
	fmt.Println("The area of square is:",b)
}
func(c circle) area()  {
	b := (c.r)*(c.r)
	fmt.Println("The area of circle is:",(math.Pi)*(b))
}