package main

import "fmt"

// type house struct{
// 	size string
// 	city string
// }

func main() {

	h := struct{
		size string
		city string
	}{
		size: "single family house",
		city : "Chicago",
	}
	fmt.Println(h)

	m := struct{
		first string
		friends map[string]int
		favDrinks []string
	}{
		first: "Kevin",
		friends: map[string]int{
			"Adam": 40,
			"Lambert": 30,
		},
		favDrinks: []string {
			"Martini",
			"Vodka",
		},

	}

	fmt.Println(m)

	for i,v := range m.friends{
		fmt.Println(i,v)
	}
	for i,v := range m.favDrinks {
		fmt.Println(i,v)
	}

}