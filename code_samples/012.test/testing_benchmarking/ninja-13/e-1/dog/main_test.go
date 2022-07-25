package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	n := Years(7)
	if n != 49 {
		fmt.Println("got", n, "want", 49)
		t.Error()
	}

}

func TestYearsTwo(t *testing.T) {

	n := YearsTwo(7)
	if n != 49 {
		fmt.Println("got", n, "want", 49)
		t.Error()
	}
}

func ExampleYears() {
	n := 7                // human year
	fmt.Println(Years(n)) // calculates god year
}

func ExampleYearsTwo() {
	a := YearsTwo(7) // human years
	fmt.Println(a)   // converst
}

func BenchmarkYears(b *testing.B) {
	// fmt.Println("dog years is: ", Years(7))
	for i := 0; i < b.N; i++ {
		Years(7)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	// fmt.Println(YearsTwo(7))
	for i := 0; i < b.N; i++ {
		YearsTwo(7)
	}
}
