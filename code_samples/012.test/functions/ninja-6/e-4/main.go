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
type k interface {
	area()
}

func main() {
	r := circle {
		r: 4,
	}
	s := square {
		a: 4,
	}
	r.area()
	s.area()
}

func(s square) area(){
	b := (s.a)*(s.a)
	fmt.Println("The area of square is:",b)
}
func(c circle) area()  {
	b := (c.r)*(c.r)
	fmt.Println("The area of circle is:",(math.Pi)*(b))
}