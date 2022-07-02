package main

import "fmt"


var t = 5

var k int

func main () {

	x := 42
	fmt.Println(x)
	x=99
	fmt.Println(x)
	y := x+42
	fmt.Println(y)

	cards := []string{"white","red","blue"}

	for _, s := range  cards { 
		fmt.Println(s, "hi")
	}

	fmt.Println(t)

	foo()

}

func foo(){
	fmt.Println("last t", t)
	k = 5
	fmt.Println(k)
	k = 10
}