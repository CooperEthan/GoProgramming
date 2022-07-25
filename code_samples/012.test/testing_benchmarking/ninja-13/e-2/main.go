package main

import (
	"example.com/v2/testing_benchmarking/ninja-13/e-1/e-2/quote"
	"example.com/v2/testing_benchmarking/ninja-13/e-1/e-2/word"
	"fmt"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
