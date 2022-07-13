package main

import (
	"fmt"
)

type person struct {
	first            string
	last             string
	ice_cream_flavor []string
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		ice_cream_flavor: []string{
			"dark",
			"pink",
			"mavi",
		},
	}

	p2 := person{
		first: "Matt",
		last:  "Damon",
		ice_cream_flavor: []string{
			"white",
			"black",
			"blue",
		},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for i, v := range m {
		fmt.Println(i)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for j, k := range v.ice_cream_flavor {
			fmt.Println(j, k)
		}
		fmt.Println("=======")
	}

	// fmt.Println(p1.first)
	// fmt.Println(p1.last)
	// for i, v := range p2.ice_cream_flavor {
	// 	fmt.Println(i,v)
	// }

}
