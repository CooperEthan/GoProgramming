package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("James")
	if s != "James" {
		// t.Error("got",s,"expected","James")
		t.Error("got", s, "expected", "James")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James"))

}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}
