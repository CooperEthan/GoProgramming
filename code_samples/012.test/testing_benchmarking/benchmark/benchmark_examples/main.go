package main

import (
	"fmt"
	"strings"

	"example.com/v2/testing_benchmarking/benchmark/benchmark_examples/mystr"
)


const s = "hey, how are you doing?"

func main() {
	xs := strings.Split(s, " ")
	for _,v := range xs {
		fmt.Println(v)
	}

	fmt.Printf("\n%s\n",mystr.Cat(xs))
	fmt.Printf("\n%s\n\n", mystr.Join(xs))
}