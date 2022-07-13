package main

import "fmt"

func main() {
	s := "Hello Playground"
	fmt.Println(s)

	k := `"Hello
	World"`
	fmt.Println(k)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\t", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println("\n ")

	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}

}
