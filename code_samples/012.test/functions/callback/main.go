package main

import "fmt"



func main() {
	fmt.Println("Hello, Playground")

	ii := []int{1,2,3,4,5,6,7,8,9,10}

	s := sum(ii...)
	fmt.Println(s)

	// fmt.Println(sum(2,3,4))

	s2 := even(sum, ii...)
	fmt.Println("total of even number is: ",s2)

	s3 := evenSum(sum,ii...)
	fmt.Println("total of odd number is: ", s3)
}

func sum( x ...int) int {

	total := 0

	for _,v := range x {
		total += v
	}
	return total
}

func even(f func(x ...int) int, y ...int) int {
	var xi []int
	for _, v := range y {
		if v%2 == 0 {
			xi = append(xi, v)
		}
	}
	total := f(xi...)
	return total
}

func evenSum(f func(x ...int)int, y ...int)int {
	var xi []int
	for _,v := range y {
		if v%2 != 0 {
			xi = append(xi, v)
		}
	}
	total := f(xi...)
	return total
}
