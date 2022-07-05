package main

import "fmt"



func main() {

	for i := 0; i < 10; i++ {
		if i%4 == 0 {
			fmt.Println(i)
		}
	}
}