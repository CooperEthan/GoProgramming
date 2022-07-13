package main

import "fmt"

func main() {

	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Miss Moneypenny"])

	v, ok := m["Barbara"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["James"]; ok {
		fmt.Println("This is the if print", v)
	}

	m["todd"] = 33

	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{4, 5, 6, 7, 8, 9, 42}
	for i, v := range xi {
		fmt.Println(i, v)
	}

	// ==============. how to delete map ==============

	delete(m, "todd")
	fmt.Println(m)
}
