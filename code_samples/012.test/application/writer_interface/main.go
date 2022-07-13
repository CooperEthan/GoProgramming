package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {

	fmt.Fprintln(os.Stdout, "Hello")
	io.WriteString(os.Stdout, "Hello")

	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	sort.Ints(xi)
	fmt.Println(xi)
	sort.Strings(xs)
	fmt.Println(xs)

}
