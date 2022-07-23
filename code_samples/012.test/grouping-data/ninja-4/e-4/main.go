package main

import "fmt"

func main() {

	a := []string{"James", "Bond", "Shaken, not stirred"}
	b := []string{"Mis", "Moneypenny", "Helloooo", "James"}

	c := [][]string{a, b}
	fmt.Println(c)

	for i, d := range c {
		fmt.Println("record: ", i)
		for j, v := range d {
			fmt.Printf("\tindex position: %v \t value: \t %v \n", j, v)
		}
	}

}
