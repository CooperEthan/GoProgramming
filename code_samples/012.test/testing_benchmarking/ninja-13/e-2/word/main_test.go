package word

import (
	"example.com/v2/testing_benchmarking/ninja-13/e-1/e-2/quote"
	"fmt"
	"testing"
)

func TestUseCount(t *testing.T) {
	// a := UseCount("one two three")
	// fmt.Println(a)
	m := UseCount("one two three three three")
	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "two":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "three":
			if v != 3 {
				t.Error("got", v, "want", 3)
			}

		}
	}
}

func TestCount(t *testing.T) {
	n := Count("one two three")
	if n != 3 {
		t.Error()
	}
}

func ExampleCount() {
	fmt.Println(Count("one two three"))
	// return 3
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
