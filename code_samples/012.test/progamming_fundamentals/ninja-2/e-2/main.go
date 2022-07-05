package main

import "fmt"


// const (
// 	a = 42
// 	b int = 42

// )

const (
	a = 2018 + iota
	b =2018 + iota
	c = 2018 + iota
	d = 2018 +iota
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// a := 42
	// fmt.Printf("%d\t%b\t%#x\t\n", a, a, a)

	// b := a << 1
	// fmt.Printf("%d\t%b\t%#x\t", b, b, b)

	// fmt.Println(a, b)

	// a := (42 == 42)
	// b := (42 <= 43)
	// c := (42 >= 43)
	// d := (42 != 43)
	// e := (42 < 43)
	// f := (42 > 43)

	// fmt.Println(a, b, c, d, e, f)

}