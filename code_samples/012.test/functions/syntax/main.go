package main

import "fmt"


func main() {

	foo()
	bar("I am from bar() function")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x,y := mouse("Ian","Fleming")
	fmt.Println(x,y)

}

// func (r reciever) identifier(parameters) return(s)) {....}

func foo() {
	fmt.Println("Hello new func")
}

func bar( s string) {
	fmt.Println(s)
}

func woo(st string) string {
	return fmt.Sprintln("Hello woo, ", st)
}

func mouse(fn string,ln string)(string,bool) {
	a := fmt.Sprintln(fn,ln, "says Hello")
	b := false
	return a,b
}