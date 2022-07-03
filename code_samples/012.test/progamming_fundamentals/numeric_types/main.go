package main

import (
	"fmt"
	"runtime"
)


var x int
var y float64

func main(){

	x = 42
	y = 42.9
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n%T\n", x, y)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOROOT())
}