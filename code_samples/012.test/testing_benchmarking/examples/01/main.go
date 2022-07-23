package main

import (
	"fmt"

	"example.com/v2/testing_benchmarking/examples/01/acdc"
)

func main() {
	fmt.Println(acdc.Sum(1, 2, 3, 4, 5))
	fmt.Println(acdc.Sum(21, 21))
}
