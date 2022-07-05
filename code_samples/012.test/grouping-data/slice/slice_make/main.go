package main

import "fmt"


func main () {

	x := make([]int, 10,12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 42
	x[9] = 999
	x = append(x, 3400, 3500, 3600)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))


	//=========== multi-dimensional slice

	b := []string{"a","b","c"}
	fmt.Println(b)
	c := []string{"ethan","cooper","ha"}
	fmt.Println(c)

	d := [][]string{b,c}
	fmt.Println(d)

}