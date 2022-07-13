// method sets
package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}

type square struct {
	r float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.r * c.r
}
func (s *square) area() float64 {
	return s.r * s.r
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {

	s := square{
		r: 5,
	}
	c := circle{
		r: 5,
	}
	info(&c)
	fmt.Printf("%T\n", c)
	info(&s)
	fmt.Printf("%T\n", s)

}
