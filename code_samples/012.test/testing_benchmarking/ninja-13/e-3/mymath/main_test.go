package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	xi := []int{1, 2, 3, 4, 5}
	a := CenteredAvg(xi)
	fmt.Println(a)

	type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		test{[]int{1, 2, 3, 4, 5}, 3},
		test{[]int{10, 20, 30, 40, 50}, 30},
		test{[]int{15, 25, 35, 45, 555}, 35},
	}
	for _, v := range tests {
		f := CenteredAvg(v.data)
		if f != v.answer {
			t.Error("got", f, "want", v.answer)
		}
	}
}

// outputs the centered avg
func ExampleCenteredAvg() {
	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(CenteredAvg(xi))
	// outputs the centered avg
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 3, 4, 5})
	}
}
