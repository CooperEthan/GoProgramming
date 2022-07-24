package main

import "fmt"

func main() {

	defer fooBar()
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(bar(b))
	fmt.Println(foo(b...))

}
func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
func fooBar() {
	defer func() {
		fmt.Println("I am ananymous func from foooBAr")
	}()
	fmt.Println("I am from fooBAr, deferred function")
}
